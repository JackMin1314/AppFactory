package main

import (
	excel "AppFactory/internal/app/HandleExcel"
	"AppFactory/internal/model"
	"AppFactory/internal/pkg/fileopr"
	log "AppFactory/pkg"
)

func main() {
	// 单列模式初始化配置
	log.InitLogger()
	logger := log.GetLogInstance()
	value := "处理excel"
	logger.Infof("初始化zap完成:%s", value)
	const excelPath = "./etc/seven2.xlsx"
	//const csvPath = "./etc/sevenSt.csv"
	const saveXlsxPath = "./etc/test2.xlsx"
	//absExcelPath, _ := filepath.Abs(excelPath)
	resultList, _ := excel.GetExcelContentList(logger, excelPath)

	err := fileopr.WriteFileData(logger, model.FILE_XLSX, saveXlsxPath, resultList)
	if err != nil {
		logger.Errorf("生成失败[%s]", err)
	}

	// for i, value := range resultList[0] {
	// 	fmt.Printf("index:%d => value:%s\n", i, value)
	// }
	// 解析到struct对象中;并写入数据库
	// WriteTableErr := excel.WriteDataToTable(logger, resultList[1:])
	// if WriteTableErr != nil {
	// 	logger.Errorf("写入数据库失败[%s]", WriteTableErr)
	// }
	logger.Info("执行结束")

}
