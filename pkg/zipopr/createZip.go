package zipopr

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

// CreatZipFile 根据srcDir路径获取所有文件, zipPathFileName为给定路径生成的zip文件
func CreatZipFile(logger *zap.SugaredLogger, zipPathFileName, srcDir string) error {
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
	os.RemoveAll(zipPathFileName)

	// 创建：zip文件
	//zipfile, err := os.Create(zipPathFileName)
	zipfile, err := os.OpenFile(zipPathFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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

// UnzipDir 从zipPathFile文件解压到dir中
func UnzipDir(logger *zap.SugaredLogger, zipPathFile, dir string) {
	absZipFile, err := filepath.Abs(zipPathFile)
	if err != nil {
		logger.Errorf("路径错误[%s]", absZipFile)
	}
	r, err := zip.OpenReader(absZipFile)
	if err != nil {
		logger.Fatalf("Open zip file failed: %s\n", err.Error())
	}
	defer r.Close()

	for _, f := range r.File {
		func() {
			path := dir + string(filepath.Separator) + f.Name
			os.MkdirAll(filepath.Dir(path), 0755)
			fDest, err := os.Create(path)
			if err != nil {
				logger.Errorf("Create failed: %s\n", err.Error())
				return
			}
			defer fDest.Close()

			fSrc, err := f.Open()
			if err != nil {
				logger.Errorf("Open failed: %s\n", err.Error())
				return
			}
			defer fSrc.Close()

			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				logger.Errorf("Copy failed: %s\n", err.Error())
				return
			}
		}()
	}
}
