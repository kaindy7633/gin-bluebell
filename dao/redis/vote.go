package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

// 投一票就加432分 ， 86400/200 --> 200张赞成票可以给帖子续一天
/**
投票的几种情况：
direction=1时，有两种情况：
	1. 之前没有投过票，现在投赞成票  -> 更新分数和投票记录 差值的绝对值：1 +432
	2. 之前投反对票，现在改投赞成票  -> 更新分数和投票记录 差值的绝对值：2 +432*2
direction=0时，有两种情况：
	1. 之前投过反对票，现在要取消投票 -> 更新分数和投票记录 差值的绝对值：1 +432
	2. 之前没有投过票，现在要取消投票 -> 更新分数和投票记录 差值的绝对值：1 -432
direction=-1时，有两种情况：
	1. 之前没有投过票，现在投反对票  -> 更新分数和投票记录 差值的绝对值：1 -432
	2. 之前投赞成票，现在改投反对票  -> 更新分数和投票记录 差值的绝对值：2 -432*2

投票的限制：
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了
	1. 到期之后将 reids 中保存的赞成票及反对票数存储到 mysql 表中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePreVote     = 432 // 每一票的分数
)

var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
)

func CreatePost(postID int64) error {
	// 加入事务操作，下面的两个操作要么都成功，要么都失败
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	_, err := pipeline.Exec()

	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票限制
	// 去redis中获取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > float64(oneWeekInSeconds) {
		return ErrVoteTimeExpire
	}

	// 2. 更新帖子的分数
	// 查询之前的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值

	// 以下操作需要放到一个pipeline事务操作中
	pipeline := rdb.TxPipeline()

	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePreVote, postID)

	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
