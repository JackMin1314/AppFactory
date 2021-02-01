package sequence

import (
	"bytes"
	"time"
)

func GenTradeID() string {
	var buf bytes.Buffer

	now := time.Now()
	buf.WriteString(now.Format("20060102"))

	id, err := GenID()
	if err != nil {
		return ""
	}
	buf.WriteString(id)
	return buf.String()
}
