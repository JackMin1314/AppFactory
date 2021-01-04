package fileopr

import (
	"AppFactory/internal/model"
	"bufio"
	"errors"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"go.uber.org/zap"
)

// WriteFileData 根据文件类型写入相应格式文件
func WriteFileData(logger *zap.SugaredLogger, fileType string, fileName string, data [][]string) error {
	switch fileType {
	case model.FILE_CSV:
		return WriteCSVFileData(logger, fileName, data)
	case model.FILE_TXT:
		return WriteTXTFileData(logger, fileName, data)
	case model.FILE_XLSX:
		return WriteXLSXFileData(logger, fileName, data)
	default:
		return errors.New("unknown file type")

	}
}

// WriteCSVFileData 写入csv格式文件
func WriteCSVFileData(logger *zap.SugaredLogger, fileName string, data [][]string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		logger.Info("打开[%s]文件失败[%s]", fileName, err)
		return err
	}
	w := bufio.NewWriter(f)
	// 不加flush会导致数据丢失
	defer w.Flush()
	i := 0
	for _, strSlice := range data {
		for _, item := range strSlice {
			//_, writeStrErr := w.WriteString(item+",")
			_, writeStrErr := w.WriteString("\uFEFF" + item + ",") // 为了适应win上wps和excel打开时候乱码问题
			if writeStrErr != nil {
				logger.Errorf("写入[%s]文件出错[%s]", fileName, writeStrErr)
				return writeStrErr
			}
		}
		i++
		_, _ = w.WriteString("\n")
	}
	logger.Infof("写入成功,总计[%s]行", i)
	return nil
}

// WriteTXTFileData 写入txt格式文件
func WriteTXTFileData(logger *zap.SugaredLogger, fileName string, data [][]string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		logger.Info("打开[%s]文件失败[%s]", fileName, err)
		return err
	}
	w := bufio.NewWriter(f)
	// 不加flush会导致数据丢失
	defer w.Flush()
	i := 0
	for _, strSlice := range data {
		for _, item := range strSlice {
			_, writeStrErr := w.WriteString(item + "\t")
			if writeStrErr != nil {
				logger.Errorf("写入[%s]文件出错[%s]", fileName, writeStrErr)
				return writeStrErr
			}
		}
		i++
		_, _ = w.WriteString("\n")
	}
	logger.Infof("写入成功,总计[%s]行", i)
	return nil
}

// WriteXLSXFileData 写入XLSX格式文件
func WriteXLSXFileData(logger *zap.SugaredLogger, fileName string, data [][]string) error {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	err := f.SetColWidth("Sheet1", "A", "Z", 15)
	if err != nil {
		logger.Errorf("设置工作表列宽度失败[%s]", err)
		return err
	}
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 获取流式写入器
	streamWriter, streamErr := f.NewStreamWriter("Sheet1")
	if streamErr != nil {
		return streamErr
	}
	colLen := len(data[0])
	rowLen := len(data)
	row := make([]interface{}, colLen)
	rowList := make([][]interface{}, len(data))

	for index, rowData := range data {
		for cindex, col := range rowData {
			row[cindex] = col
		}
		rowList[index] = row
		cell, _ := excelize.CoordinatesToCellName(1, index+1) // col: 1, row: 1 => A1
		if err := streamWriter.SetRow(cell, rowList[index]); err != nil {
			logger.Errorf("写入excel失败[%s]", err)
			return err
		}
	}
	if err := streamWriter.Flush(); err != nil {
		logger.Errorf("flush excel失败[%s]", err)
		return err
	}
	// 设置样式
	styleID, err := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"alignment":{"horizontal":"center"}}`)
	cell, _ := excelize.CoordinatesToCellName(colLen, rowLen)
	_ = f.SetCellStyle("Sheet1", "A1", cell, styleID)
	if err := f.SaveAs(fileName); err != nil {
		logger.Errorf("保存excel失败[%s]", err)
		return err
	}
	return nil
}
