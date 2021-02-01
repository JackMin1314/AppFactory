package redisopr

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisCli struct {
	redis.Conn
}

func NewRedis() *RedisCli {
	return &RedisCli{Conn: pool.Get()}
}

func (c *RedisCli) Close() {
	c.Conn.Close()
}

func (c *RedisCli) OprExist(key string) (bool, error) {
	return redis.Bool(c.Conn.Do("EXISTS", key))
}
func (c *RedisCli) OprExpire(key string, seconds int) error {
	_, err := c.Conn.Do("EXPIRE", key, seconds)
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCli) OprExpireMillSeconds(key string, millseconds int) error {
	_, err := c.Conn.Do("PEXPIRE", key, millseconds)
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCli) OprDel(key string) error {
	_, err := c.Conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCli) OprSet(key string, val interface{}) error {
	_, err := c.Conn.Do("SET", key, val)
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCli) OprGet(key string) ([]byte, error) {
	val, err := redis.Bytes(c.Conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *RedisCli) OprListAdd(key string, val interface{}) error {
	_, err := c.Conn.Do("RPUSH", key, val)
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCli) OprListGet(key string, indexs ...int) ([][]byte, error) {
	begin := 0
	end := -1
	if len(indexs) == 2 {
		begin = indexs[0]
		end = indexs[1]
	}
	if ok, err := c.OprExist(key); err != nil || !ok {
		return nil, fmt.Errorf("%v,%v", err, ok)
	}
	return redis.ByteSlices(c.Conn.Do("LRANGE", key, begin, end))
}

func (c *RedisCli) OprZSetAdd(key string, proi int, val interface{}) error {
	_, err := c.Conn.Do("ZADD", key, proi, val)
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCli) OprZSetGet(key string, indexs ...int) ([][]byte, error) {
	begin := 0
	end := -1
	if len(indexs) == 2 {
		begin = indexs[0]
		end = indexs[1]
	}
	return redis.ByteSlices(c.Conn.Do("ZRANGE", key, begin, end))
}
