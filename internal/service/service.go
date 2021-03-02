package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

/*
/internal/service 实现了 api 定义的服务层
类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
service中作为协调任务和分配工作的应用层，直接用的是biz中的结构体，而不是biz中的接口（结构体有方法提现任务，其次也是为了易于扩展性和维护）进行了封装；
提供了业务的接口流程（链路追踪也是在这块有体现）
*/
