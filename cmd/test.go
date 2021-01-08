package main

import (
	"AppFactory/internal/dao"
	"AppFactory/internal/model"
	log "AppFactory/pkg"
	"fmt"
)

func main() {
	log.InitLogger()
	logger := log.GetLogInstance()
	value := "处理excel"
	logger.Infof("初始化zap完成:%s", value)
	var IEO model.IDBOpt
	// QEA := dao.QueryExamAll
	QEA := dao.QueryExamMain
	IEO = QEA
	qeaResults, err := IEO.FindTableAll(logger, "step_rank>100 and step_rank<200") // ["step_rank", "27"],["step_rank>100"],["step_rank>",100] , ["step_rank>100 and step_rank<200"]
	if err != nil {
		logger.Errorf("查询失败[%s]", err)
	}
	qeaResultList := qeaResults.([]*model.QueryExamMain)
	for index, item := range qeaResultList {
		fmt.Println(index, "=>", *item)
	}
	err = QEA.SoftDeletRecord(logger, "is_deleted='1'", "upload_date", "20210106")
	fmt.Println(err)

}
