package excel

import (
	"errors"

	"AppFactory/internal/dao"
	"AppFactory/internal/model"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"go.uber.org/zap"
)

const sheet = "Sheet1"

//

// GetExcelContentList 是用来获取包含excel表格列数据列表
func GetExcelContentList(logger *zap.SugaredLogger, excelPath string) ([][]string, error) {
	f, err := excelize.OpenFile(excelPath)

	if err != nil {
		logger.Errorf("打开Excel表格[%s]失败[%s]", excelPath, err)
		return nil, err
	}

	// 获取excel表格所有表
	sheetNameList := f.GetSheetList()
	// 默认处理第一张表
	var sheetName = ""
	if len(sheetNameList) >= 1 {
		sheetName = sheetNameList[0]
	} else {
		return nil, errors.New("excel表格内容为空")
	}

	// 获取表格所有行
	rows, getRowsErr := f.GetRows(sheetName)
	if getRowsErr != nil {
		logger.Errorf("按行获取excel表格内容失败[%s]", getRowsErr)
		return nil, getRowsErr
	}
	if len(rows[0]) <= 1 {
		logger.Error("excel解析失败，请检查内容是否含有标题")
		return nil, errors.New("excel解析失败，请检查内容是否含有标题")
	}
	excelRows := make([][]string, 0)
	for rindex, row := range rows {
		rowLen := len(row)
		if rowLen == 0 {
			logger.Infof("解析到第[%d]行无数据内容，跳过")
			break
		}
		emptyNum := 1
		for _, col := range row {

			if col == "" && emptyNum <= rowLen {
				emptyNum++
				//excelRows[rindex][cindex] = "0"
				continue
			}
		}

		if emptyNum > rowLen {
			logger.Infof("第[%d]行读取到空行数据，停止解析", rindex)
			excelRows = excelRows[0:rindex]
			return excelRows, nil
		}
		excelRows = append(excelRows, rows[rindex])
	}

	return excelRows, nil
}

// WriteDataToTable 根据数据的不同类型，选择插入到特定表中
func WriteDataToTable(logger *zap.SugaredLogger, datasAll [][]string) error {

	db := g.DB()
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		logger.Errorf("开启事务失败[%s]", err)
		return err
	}
	// len返回是下标为1的长度计数
	lenNum := len(datasAll[0])
	switch lenNum {
	// all
	case 12:
		logger.Info("表格列字段为[12]")
		return InsertQEATable(logger, tx, datasAll)

	case 10:
		logger.Info("表格列字段为[10]")
		return InsertQEMTable(logger, tx, datasAll)
	default:
		logger.Infof("表格列字段为[%d]", lenNum)
		return errors.New("该列长度未进行判断")
	}
}

// InsertQEATable 带事务写入query_exam_all数据库逻辑
func InsertQEATable(logger *zap.SugaredLogger, tx *gdb.TX, datasAll [][]string) error {
	QueryEADaoTx := dao.QueryExamAll.TX(tx)
	QEAList := make([]model.QueryExamAll, 0)
	var qea model.QueryExamAll
	for _, rowList := range datasAll {
		qea.ExamNum = rowList[0]
		qea.StudentName = rowList[1]
		qea.ClassName = rowList[2]
		qea.ChineseScore = rowList[3]
		qea.MathScore = rowList[4]
		qea.EnglishScore = rowList[5]
		qea.PoliticsScore = rowList[6]
		qea.HistoryScore = rowList[7]
		qea.TotalScore = rowList[8]
		qea.ClassRate = rowList[9]
		qea.SchoolRate = rowList[10]
		qea.StepRank = rowList[11]

		QEAList = append(QEAList, qea)
	}

	sqlResult, err := QueryEADaoTx.Insert(QEAList)
	if err != nil {
		logger.Errorf("插入数据库表[%s]数据失败[%s]", dao.QueryExamAll.Table, err)
		tx.Rollback()
	}
	tx.Commit()
	if id, _ := sqlResult.RowsAffected(); id != 0 {
		logger.Infof("插入数据库表[%s]数据成功id[%d]", dao.QueryExamAll.Table, id)
	}
	return nil
}

// InsertQEMTable 带事务写入query_exam_main数据库逻辑
func InsertQEMTable(logger *zap.SugaredLogger, tx *gdb.TX, datasAll [][]string) error {

	QueryEMDaoTx := dao.QueryExamMain.TX(tx)
	QEMList := make([]model.QueryExamMain, 0)
	var qea model.QueryExamMain
	for _, rowList := range datasAll {
		qea.ExamNum = rowList[0]
		qea.StudentName = rowList[1]
		qea.ClassName = rowList[2]
		qea.ChineseScore = rowList[3]
		qea.MathScore = rowList[4]
		qea.EnglishScore = rowList[5]
		qea.TotalScore = rowList[6]
		qea.ClassRate = rowList[7]
		qea.SchoolRate = rowList[8]
		qea.StepRank = rowList[9]

		QEMList = append(QEMList, qea)
	}

	sqlResult, err := QueryEMDaoTx.Insert(QEMList)
	if err != nil {
		logger.Errorf("插入数据库表[%s]数据失败[%s]", dao.QueryExamMain.Table, err)
		tx.Rollback()
	}
	tx.Commit()
	if id, _ := sqlResult.RowsAffected(); id != 0 {
		logger.Infof("插入数据库表[%s]数据成功id[%d]", dao.QueryExamMain.Table, id)
	}
	return nil
}
