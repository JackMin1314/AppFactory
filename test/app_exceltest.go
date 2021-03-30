package main

import (
	"errors"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"go.uber.org/zap"
)

const sheet = "Sheet1"

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
