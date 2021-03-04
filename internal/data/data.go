package data

import (
	"AppFactory/pkg/config"
	"AppFactory/pkg/log"
	"time"

	"AppFactory/internal/data/ent"

	"github.com/go-redis/redis/extra/redisotel"
	redis "github.com/go-redis/redis/v8"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,NewAppExcelImplRepo)

// Data contains db and redis
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(c *config.ConfigYaml, logger *log.ZapLog) (*Data, error) {
	client, err := ent.Open(c.Data.DataBase.Driver, c.Data.DataBase.Source)
	if err != nil {
		logger.Errorf("failed opening connection to sqlite: %v", err)
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Data.Redis.Addr,
		Password:     c.Data.Redis.Password,
		DB:           c.Data.Redis.Db,
		DialTimeout:  time.Duration(c.Data.Redis.DialTimeout),
		WriteTimeout: time.Duration(c.Data.Redis.WriteTimeout),
		ReadTimeout:  time.Duration(c.Data.Redis.ReadTimeout),
	})
	rdb.AddHook(redisotel.TracingHook{})
	return &Data{
		db:  client,
		rdb: rdb,
	}, nil
}

/*
/internal/data 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。
浅显理解 DDD 的 infra层，data 偏重业务的含义，它所要做的是将领域对象重新拿出来（不同于dao处理数据），我们去掉了 DDD 的 infra层。
data里面封装基本的db和rdb外，还实现了biz中定义的一些接口（在biz中会被组合在一起放在结构体里），通过NewXX返回这个biz中的**接口**，便于给main中的wire
*/
