package sequence

import (
	"strconv"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	t, _ := time.Parse("20060102", "20200101")
	settings := sonyflake.Settings{
		StartTime: t,
	}
	sf = sonyflake.NewSonyflake(settings)
}

func GenID() (string, error) {
	id, err := sf.NextID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 10), nil
}
