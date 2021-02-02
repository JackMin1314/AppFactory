package redisopr

import "strings"

type RedisKey struct {
	prefix string
}

func (c *RedisKey) Key(tp, id string) string {
	str := strings.Builder{}
	str.WriteString(c.prefix)
	str.WriteString("::")
	str.WriteString(tp)
	str.WriteString("::")
	str.WriteString(id)
	return str.String()
}

//全局封装函数
func Key(tp, id string) string {
	return redisKey.Key(tp, id)
}
