// package main

// import (
// 	"AppFactory/internal/dao"
// 	"AppFactory/internal/model"
// 	log "AppFactory/pkg"
// 	"fmt"
// )

// func main() {
// 	log.InitLogger()
// 	logger := log.GetLogInstance()
// 	value := "处理excel"
// 	logger.Infof("初始化zap完成:%s", value)
// 	var IEO model.IDBOpt
// 	// QEA := dao.QueryExamAll
// 	QEA := dao.QueryExamMain
// 	IEO = QEA
// 	qeaResults, err := IEO.FindTableAll(logger, "step_rank>100 and step_rank<200") // ["step_rank", "27"],["step_rank>100"],["step_rank>",100] , ["step_rank>100 and step_rank<200"]
// 	if err != nil {
// 		logger.Errorf("查询失败[%s]", err)
// 	}
// 	qeaResultList := qeaResults.([]*model.QueryExamMain)
// 	for index, item := range qeaResultList {
// 		fmt.Println(index, "=>", *item)
// 	}
// 	err = QEA.SoftDeletRecord(logger, "is_deleted='1'", "upload_date", "20210106")
// 	fmt.Println(err)

// }
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	
	str1 := "hello"
	str2 := "world"
	fmt.Println(FMTSPrintf(str1, str2))
}

func FMTSPrintf(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func AddOper(a, b string) string {
	return a + b
}

func StrJoin(a, b string) string {
	return strings.Join([]string{a, b}, "")
}

func ByteBuffer(a, b string) string {
	//定义Buffer类型
	var bt bytes.Buffer
	// 向bt中写入字符串
	bt.WriteString(a)
	bt.WriteString(b)
	//获得拼接后的字符串
	s3 := bt.String()
	return s3

}

func StringsBuilder(a, b string) string {
	var build strings.Builder
	build.WriteString(a)
	build.WriteString(b)
	s3 := build.String()
	return s3
}
