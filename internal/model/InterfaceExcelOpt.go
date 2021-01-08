package model

import (
	"go.uber.org/zap"
)

// IDBOpt 是处理excel通用的方法集合
type IDBOpt interface {
	InsertTableAll(*zap.SugaredLogger, [][]string) error
	FindTableOne(*zap.SugaredLogger, ...interface{}) (interface{}, error)
	FindTableAll(*zap.SugaredLogger, ...interface{}) (interface{}, error)
	SoftDeletRecord(*zap.SugaredLogger, ...interface{}) error
}
