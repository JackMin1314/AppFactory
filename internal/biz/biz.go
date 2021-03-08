package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAppExcelUsecase)
/*
/internal/biz 业务逻辑的组装层
类似 DDD 的 domain 层，data 类似 DDD 的 repo，repo 接口在这里定义，使用依赖倒置的原则（业务层的抽象）
biz中组合接口的结构体，出于屏蔽底层细节和对外向service（Application层）提供方法的考虑，
对其内接口内定义的每一个方法，都提供了方法实现对接口内部函数的调用（接口变量调用本身的方法函数），同时通过NewXX返回biz中该 结构体，便于给main中wire。
*/