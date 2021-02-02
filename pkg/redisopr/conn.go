package redisopr

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"


	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

var (
	redisOnce sync.Once
	pool      *redis.Pool
	config    *redisConfig = new(redisConfig)
	redisKey  *RedisKey
)

type redisConfig struct {
	Redis struct {
		PreFix   string `yaml:"PREFIX"`
		Addr     string `yaml:"ADDR"`
		Password string `yaml:"PASSWORD"`

		MaxActive int `yaml:"MAXACTIVE"`
		MaxIdle   int `yaml:"MAXIDLE"`

		IdleTimeout  time.Duration `yaml:"IDLETIMEOUT"` //单位为秒
		ConnTimeout  time.Duration `yaml:"CONNTIMEOUT"` //单位为秒
		WriteTimeout time.Duration `yaml:"WRTIMEOUT"`   //单位为秒
		ReadTimeout  time.Duration `yaml:"RDTIMEOUT"`   //单位为秒
	} `yaml:"REDIS"`
}

const defaultMaxActiveConn = 1000
const defaultMaxIdleConn = 500
const defaultIdleTimeout = 300 * time.Minute
const defaultConnTimeout = 60 * time.Second
const defaultWriteTimeout = 300 * time.Second
const defaultReadTimeout = 300 * time.Second

func Init(filename string) error {
	fullname, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	// TODO:解析加载配置文件,后面进行优化处理
	err = yamlcfg.LoadConfig(fullname, config)
	if err != nil {
		return fmt.Errorf("加载配置文件[%s]失败[%s]", fullname, err)
	}

	if config.Redis.MaxActive == 0 {
		config.Redis.MaxActive = defaultMaxActiveConn
	}
	if config.Redis.MaxIdle == 0 {
		config.Redis.MaxIdle = defaultMaxIdleConn
	}

	if config.Redis.IdleTimeout == 0 {
		config.Redis.IdleTimeout = defaultIdleTimeout
	} else {
		config.Redis.IdleTimeout *= time.Second
	}

	if config.Redis.ConnTimeout == 0 {
		config.Redis.ConnTimeout = defaultConnTimeout
	} else {
		config.Redis.ConnTimeout *= time.Second
	}

	if config.Redis.WriteTimeout == 0 {
		config.Redis.WriteTimeout = defaultWriteTimeout
	} else {
		config.Redis.WriteTimeout *= time.Second

	}

	if config.Redis.ReadTimeout == 0 {
		config.Redis.ReadTimeout = defaultReadTimeout
	} else {
		config.Redis.ReadTimeout *= time.Second
	}

	err = errors.Errorf("已经初始化,不再重复初始化")
	//连接Redis
	fmt.Printf("redisopr连接地址:[%+v]\n", config.Redis)

	fmt.Println("redisOpr连接")
	redisOnce.Do(func() {
		//设置KEY前缀
		redisKey = new(RedisKey)
		redisKey.prefix = config.Redis.PreFix

		pool = &redis.Pool{
			MaxIdle:     config.Redis.MaxIdle,
			IdleTimeout: config.Redis.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", config.Redis.Addr,
					redis.DialPassword(config.Redis.Password),
					redis.DialConnectTimeout(config.Redis.ConnTimeout),
					redis.DialReadTimeout(config.Redis.ReadTimeout),
					redis.DialWriteTimeout(config.Redis.WriteTimeout))
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, terr := c.Do("PING")
				return terr
			},
			MaxActive: config.Redis.MaxActive,
			Wait:      true,
		}
		c := pool.Get()
		defer c.Close()

		_, err = c.Do("PING")
		return
	})
	return err
}
