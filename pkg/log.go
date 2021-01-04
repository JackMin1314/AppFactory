package log

import (
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugarLogger *zap.SugaredLogger

var once sync.Once

// 日志初始化
func InitLogger() {

	once.Do(func() {
		cfg := zap.NewProductionEncoderConfig()
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		writeSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel),
			zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
		logger := zap.New(core, zap.AddCaller())
		defer logger.Sync()
		sugarLogger = logger.Sugar()
		sugarLogger.Info("first init logger success")
		return
	})
}

// 日志写入文件
func getLogWriter() zapcore.WriteSyncer {
	folderPath := "./logs"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	folderPath = filepath.Join(folderPath, "Record.log")
	//_, err := os.Create(folderPath)
	_, err := os.OpenFile(folderPath,os.O_APPEND|os.O_CREATE, 0666)// os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)会覆盖或重新创建
	if err != nil {
		return nil
	}
	// 新增日志切割和归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   folderPath, // 日志文件的位置
		MaxSize:    10,         // 在进行切割之前，日志文件的最大大小
		MaxBackups: 30,         // 保留旧文件的最大个数
		MaxAge:     7,          // 保留旧文件的最大天数
		Compress:   false,      // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志文件使用大写字母记录日志级别
	return zapcore.NewConsoleEncoder(encodeConfig)
}

func GetLogInstance() *zap.SugaredLogger {
	if sugarLogger != nil {
		return sugarLogger
	}
	return nil
}
