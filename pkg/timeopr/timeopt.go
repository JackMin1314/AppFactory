package timeopr

import (
	"context"
	"fmt"
	"time"
)

// GetBeforeOneDay 用于获取指定日期的前一天
func GetBeforeOneDay(dateStr string) string {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	st := dateStr + " 23:00:00"
	t, _ := time.ParseInLocation("20060102 15:04:05", st, LOC)

	//fmt.Println(t.Unix())
	//tt := time.Unix(t.Unix(), 0)
	//fmt.Println(tt.Format("2006-01-02 15:04:05"))

	yesterday := t.AddDate(0, 0, -1)
	yesterdayDate := yesterday.Format("20060102 15:04:05")
	return yesterdayDate[:8]
}

// GetDateTimeFormat 获取时间年月日、月、日等
func GetDateTimeFormat() {
	allDate := time.Now().Format("20060102")
	YearMD := time.Now().Format("060102")
	YearM := time.Now().Format("200601")
	MonthDay := time.Now().Format("0102")
	fmt.Println(allDate, YearMD, YearM, MonthDay)
}

// 超时小于给定的时间，则用超时的时间创建
func ShrinkDeadline(ctx context.Context, timeout time.Duration) (context.Context, func()) {
	if deadline, ok := ctx.Deadline(); ok {
		leftTime := time.Until(deadline)
		if leftTime < timeout {
			timeout = leftTime
		}
	}

	return context.WithDeadline(ctx, time.Now().Add(timeout))
}
