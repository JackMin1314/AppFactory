package main

import (
	"fmt"
	// "log"
	"os"
	"os/signal"
	"syscall"

	log "AppFactory/pkg"

	"github.com/gin-gonic/gin"
)

/* 需要导入log
func main() {

	// check if exist notifyLog.log
	_, errfile := os.Stat("notifyLog.log")
	if errfile != nil {
		if os.IsNotExist(errfile) {
			os.Create("notifyLog.log")
		}
	}

	logFile, err := os.OpenFile("notifyLog.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	gin.SetMode("release")
	fmt.Println("server info ==>> IP_PORT:[0.0.0.0:8666] Router:[/notify] Method:[GET POST] Log:[./notifyLog.log]")
	router := gin.Default()
	router.GET("/notify", func(context *gin.Context) {
		log.Println(">>>> receive notify request <<<<")
		notifyMsg, rawerr := context.GetRawData()
		if rawerr != nil {
			log.Printf("读取请求报文失败[%s]\n", rawerr)

			context.String(200, "ok")
			return
		}
		log.Printf("收到请求报文[%s]\n", string(notifyMsg))
		context.String(200, "ok")
	})

	router.POST("/notify", func(context *gin.Context) {
		log.Println(">>>> receive notify request <<<<")
		notifyMsg, rawerr := context.GetRawData()
		if rawerr != nil {
			log.Printf("读取请求报文失败[%s]\n", rawerr)
			context.String(200, "ok")
			return
		}
		log.Printf("收到请求报文[%s]\n", string(notifyMsg))
		context.String(200, "ok")
	})

	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				GracefullExit()
			default:
				fmt.Println("receive other signal", s)
			}
		}
	}()

	// 指定地址和端口号
	router.Run("0.0.0.0:8666")
}
*/
func main() {
	log.InitLogger()
	logger := log.GetLogInstance()
	gin.SetMode("release")
	fmt.Println("server info ==>> IP_PORT:[0.0.0.0:8666] Router:[/notify] Method:[GET POST] Log:[./notifyLog.log]")
	router := gin.Default()
	router.GET("/notify", func(context *gin.Context) {
		logger.Infof(">>>> receive notify request <<<<")
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
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

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
