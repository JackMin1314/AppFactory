package log

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	// "github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/os/gcfg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugarLogger *zap.SugaredLogger

var once sync.Once

// InitLogger 日志初始化
func InitLogger() {

	once.Do(func() {
		cfg := zap.NewProductionEncoderConfig()
		// cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}
		writeSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel),
			zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 终端和日志都记录
		)
		logger := zap.New(core, zap.AddCaller())
		defer logger.Sync()
		sugarLogger = logger.Sugar()
		sugarLogger.Info("first init logger success")
		return
	})
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	} //zapcore.ISO8601TimeEncoder   // 修改时间
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志文件使用大写字母记录日志级别 //zapcore.CapitalColorLevelEncoder // 按级别取值颜色

	return zapcore.NewConsoleEncoder(encodeConfig)
}

// 日志写入文件
func getLogWriter() zapcore.WriteSyncer {
	config := gcfg.Instance("config.toml")
	if config == nil {
		panic("config.toml file is not correct or not in config directort")
	}
	config.GetJson("log.config")
	folderPath := config.GetString("log.config.filePath")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	// 根据配置文件配置的日志路径和文件名
	folderPath = filepath.Join(folderPath, config.GetString("log.config.fileName"))
	//_, err := os.Create(folderPath)
	_, err := os.OpenFile(folderPath, os.O_APPEND|os.O_CREATE, 0666) // os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)会覆盖或重新创建
	if err != nil {
		return nil
	}
	// 新增日志切割和归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   folderPath,                  // 日志文件的位置
		MaxSize:    config.GetInt("MaxSize"),    // 在进行切割之前，日志文件的最大大小
		MaxBackups: config.GetInt("MaxBackups"), // 保留旧文件的最大个数
		MaxAge:     config.GetInt("MaxAge"),     // 保留旧文件的最大天数
		Compress:   config.GetBool("Compress"),  // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func GetLogInstance() *zap.SugaredLogger {
	if sugarLogger != nil {
		return sugarLogger
	}
	return nil
}
