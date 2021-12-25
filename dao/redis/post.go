package redis

import "gin-bluebell/models"

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从 redis 中获取 id
	// 根据用户请求中携带的 order 参数确定要查询的 redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 确定查询的索引起始点
	start := (p.Page - 1) * p.PageSize
	end := start + p.PageSize - 1
	// 查询
	return rdb.ZRevRange(key, start, end).Result()
}
