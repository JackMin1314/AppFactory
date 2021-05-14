package main

import (
	config "AppFactory/pkg/config"
	log "AppFactory/pkg/log"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	var confName string
	flag.StringVar(&confName, "c", "", "the explicit yarm configuration file path")
	flag.Parse()
	// 组件初始化
	cfgYaml := config.InitConfigYaml(confName)
	// 日志初始化
	log.InitLogger(cfgYaml)
	logger := log.GetLogInstance()
	gin.SetMode("release")
	fmt.Println("server info ==>> IP_PORT:[0.0.0.0:8666] Router:[/notify] Method:[GET POST] Log:[./notifyLog.log]")
	router := gin.Default()
	router.GET("/notify", func(context *gin.Context) {
		logger.Infof(">>>> receive notify request <<<<")
		logger.Infof("content-type:[%s]", context.ContentType())
		logger.Infof("remote-ip:[%s]", context.Request.RemoteAddr)
		notifyMsg, rawerr := context.GetRawData()
		if rawerr != nil {
			logger.Infof("读取请求报文失败[%s]", rawerr)

			// context.String(200, "SUCCESS")
			context.String(200, "ok")
			return
		}
		logger.Infof("收到请求报文[%s]", string(notifyMsg))
		context.String(200, "ok")
	})

	router.POST("/notify", func(context *gin.Context) {
		logger.Infof(">>>> receive notify request <<<<")
		logger.Infof("content-type:[%s]", context.ContentType())
		logger.Infof("remote-ip:[%s]", context.Request.RemoteAddr)
		notifyMsg, rawerr := context.GetRawData()
		if rawerr != nil {
			logger.Infof("读取请求报文失败[%s]", rawerr)
			context.String(200, "ok")
			return
		}
		logger.Infof("收到请求报文[%s]", string(notifyMsg))
		context.String(200, "ok")
	})

	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				logger.Infof("receive signal")
				GracefullExit()
			default:
				fmt.Println("receive other signal", s)
			}
		}
	}()

	// 指定地址和端口号
	router.Run("0.0.0.0:8666")

}
func GracefullExit() {
	os.Exit(0)
}
