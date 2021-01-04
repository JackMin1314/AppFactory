package pkg

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

func CreatZipFile(logger *zap.SugaredLogger, srcDir string, zipFileName string) error {
	dir, err := ioutil.ReadDir(srcDir)
	if err != nil {
		logger.Errorf("读取路径出错[%s]", err)
		return err
	}
	if len(dir) == 0 {
		logger.Infof("路径[%s]为空", srcDir)
		return nil
	}
	// 预防：旧文件无法覆盖
	os.RemoveAll(zipFileName)

	// 创建：zip文件
	//zipfile, err := os.Create(zipFileName)
	zipfile, err := os.OpenFile(zipFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer zipfile.Close()
	if err != nil {
		logger.Errorf("创建zip文件出错[%s]", err)
		return err
	}
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			logger.Errorf("获取文件头信息出错[%s]", err)
			return err
		}

		header.Name = path
		// for win
		// header.Name = strings.TrimPrefix(path, srcDir+`\`)
		// for linux
		header.Name = strings.TrimPrefix(path, srcDir+`/`)

		// 判断：文件是不是文件夹
		//if info.IsDir() {
		//	header.Name += `/`
		//} else {
		//	// 设置：zip的文件压缩算法
		//	header.Method = zip.Deflate
		//}

		// 设置：zip的文件压缩算法
		header.Method = zip.Deflate
		// 创建：压缩包头部信息
		writer, err := archive.CreateHeader(header)
		if err != nil {
			logger.Errorf("创建压缩包头部信息出错[%s]", err)
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				logger.Errorf("打开文件出错[%s]", err)
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}

		return err
	})
	return nil
}
