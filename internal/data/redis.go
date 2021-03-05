package data

import (
	"AppFactory/internal/biz"
	"context"
	"encoding/json"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func FormatRedisKey(stu *biz.AppExcel) string {
	if stu == nil {
		return ""
	}
	return stu.ExamNum
}

func (data *Data) RDPing(ctx context.Context) error {
	if _, err := data.rdb.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}

func (data *Data) RDSetStruct(ctx context.Context, key string, value interface{}) error {
	stData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return data.rdb.Set(ctx, key, stData, time.Duration(30)*time.Second).Err()
}

func (data *Data) RDGetStruct(ctx context.Context, key string) ([]byte, error) {
	valByte, err := data.rdb.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return valByte, nil
		}
		return nil, err
	}
	return valByte, nil
}
