package zipmake

import (
	"io/ioutil"
	"os"
	"path/filepath"

	pkg "AppFactory/internal/pkg/zip"

	"go.uber.org/zap"
)

func GenSettZipFiles(logger *zap.SugaredLogger) error {

	zipPathStr := "/data/reportFiles"
	SettDate := "20201215"
	settleFilePath := "/data/settfiles"
	_, err := os.Stat(filepath.Join(zipPathStr, SettDate))
	if os.IsNotExist(err) {
		logger.Infof("开始创建报表压缩文件夹根路径[%s]", zipPathStr)
		// 必须分成两步：先创建文件夹、再修改权限
		if err = os.MkdirAll(filepath.Join(zipPathStr, SettDate), 0777); err != nil {
			logger.Errorf("创建报表压缩文件夹路径[%s]失败,", zipPathStr)
			return err
		}
		//0777也可以os.ModePerm;默认权限修改777失败时，不return
		if err = os.Chmod(zipPathStr, 0777); err != nil {
			logger.Errorf("修改文件夹路径[%s]权限为0777失败,", zipPathStr)
		}
		logger.Infof("创建报表文件夹路径[%s]成功", zipPathStr)
	}
	absZipPath, _ := filepath.Abs(zipPathStr)
	absZipDatePath := filepath.Join(absZipPath, SettDate)

	// 生成总报表压缩文件
	TotalZipFile := filepath.Join(absZipDatePath, SettDate+"_all.zip")
	logger.Infof("开始生成总报表压缩文件[%s]", SettDate+"_all.zip")
	srcSettFilePath := filepath.Join(settleFilePath, SettDate)
	totalZipFileErr := pkg.CreatZipFile(logger, srcSettFilePath, TotalZipFile)
	if totalZipFileErr != nil {
		logger.Errorf("生成总报表压缩文件失败[%s]", totalZipFileErr)
		return totalZipFileErr
	}
	logger.Infof("总报表压缩文件[%s]生成成功", SettDate+"_all.zip")

	// 生成商户报表压缩文件和门店报表压缩文件
	dirlist, readErr := ioutil.ReadDir(filepath.Join(settleFilePath, SettDate))
	if readErr != nil {
		return readErr
	}

	// TODO: bad code
	logger.Infof("开始生成[%s]日期商户、门店报表压缩文件", SettDate)
	for _, MchtDirItem := range dirlist {
		// 处理当前商户
		if MchtDirItem.IsDir() {
			logger.Infof("开始处理商户[%s]报表压缩文件", MchtDirItem.Name())
			// 生成商户报表压缩文件
			mchtSrcPath := filepath.Join(settleFilePath, SettDate, MchtDirItem.Name())
			MchtZipFile := filepath.Join(absZipDatePath, MchtDirItem.Name()+".zip")
			if genMchtZipErr := pkg.CreatZipFile(logger, mchtSrcPath, MchtZipFile); genMchtZipErr != nil {
				logger.Errorf("生成商户[%s]报表压缩文件失败[%s]", MchtDirItem.Name(), genMchtZipErr)
				return genMchtZipErr
			}
			// 处理商户下的门店报表压缩文件
			shoplist, readErr := ioutil.ReadDir(mchtSrcPath)
			if readErr != nil {
				return readErr
			}

			for _, ShopDirItem := range shoplist {
				// 生成门店报表压缩文件
				shopSrcPath := filepath.Join(mchtSrcPath, ShopDirItem.Name())
				ShopZipFile := filepath.Join(absZipDatePath, ShopDirItem.Name()+".zip")
				if ShopDirItem.IsDir() {
					if genShopZipErr := pkg.CreatZipFile(logger, shopSrcPath, ShopZipFile); genShopZipErr != nil {
						logger.Errorf("生成商户[%s]的门店[%s]报表压缩文件失败[%s]", MchtDirItem.Name(), ShopDirItem.Name(), genShopZipErr)
						return genShopZipErr
					}
				}

			}

		} // end  MchtDirItem.IsDir()

	}
	logger.Infof("[%s]日期商户、门店报表压缩文件生成成功", SettDate)

	return nil
}
