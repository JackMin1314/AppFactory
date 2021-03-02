package data

import (
	"AppFactory/internal/conf"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data) (*Data, error) {
	return &Data{}, nil
}

/*
/internal/data 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。
浅显理解 DDD 的 infra层，data 偏重业务的含义，它所要做的是将领域对象重新拿出来（不同于dao处理数据），我们去掉了 DDD 的 infra层。
data里面封装基本的db和rdb外，还实现了biz中定义的一些接口（在biz中会被组合在一起放在结构体里），通过NewXX返回这个biz中的**接口**，便于给main中的wire
*/
