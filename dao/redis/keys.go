package redis

// redis key
// redis key 注意使用命名空间的方式
const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"   // zset;贴子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;帖子及投票分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;参数是 post id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
