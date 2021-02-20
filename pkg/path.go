package tools

import (
	"os"
	"strings"
	"time"
)

var gCurrDir string = ""

func ConfigInit() error {
	var err error
	gCurrDir, err = os.Getwd()
	if err != nil {
		return err
	}
	const APPDIR = "APPHOME"
	appHome := os.Getenv(APPDIR)
	if appHome == "" {
		appHome = gCurrDir
	}

	err = os.Chdir(appHome)
	if err != nil {
		return err
	}

	return nil
}

func ConfigEnd() error {
	var err error
	err = os.Chdir(gCurrDir)
	if err != nil {
		return err
	}
	return nil
}

// 格式化时间
func ParseFormat(t time.Time, orig string) string {
	resu := strings.Replace(orig, "YYYYMMDD", "20060102", -1)
	resu = strings.Replace(resu, "HHMMSS", "150405", -1)
	return t.Format(resu)
}

// 批量创建文件路径
func MkdirFilePath(savePath []string) error {
	for _, value := range savePath {
		if value == "" {
			continue
		}
		_, err := os.Stat(value)
		if os.IsNotExist(err) {
			// 必须分成两步：先创建文件夹、再修改权限
			if err = os.MkdirAll(value, 0777); err != nil {
				return err
			}
			//0777也可以os.ModePerm;默认权限修改777失败时，不return
			// if err = os.Chmod(value, 0777); err != nil {
			// }
			_ = os.Chmod(value, 0777)
		}
	}
	return nil
}
