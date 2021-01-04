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

func ParseFormat(t time.Time, orig string) string {
	resu := strings.Replace(orig, "YYYYMMDD", "20060102", -1)
	resu = strings.Replace(resu, "HHMMSS", "150405", -1)
	return t.Format(resu)
}
