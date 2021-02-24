package log

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"AppFactory/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugarLogger *zap.SugaredLogger

var once sync.Once

// InitLogger 日志初始化
func InitLogger(cfgYml *config.ConfigYaml) {

	once.Do(func() {
		cfg := zap.NewProductionEncoderConfig()
		// cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}
		writeSyncer := getLogWriter(cfgYml)
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
func getLogWriter(cfg *config.ConfigYaml) zapcore.WriteSyncer {
	if cfg == nil {
		panic("configYaml is not initial or the file is not in configs directory")
	}
	if cfg.Log.FileFolder == "" {
		panic("log.filePath is not correct in config yaml")
	} else {
		if _, err := os.Stat(cfg.Log.FileFolder); os.IsNotExist(err) {
			// 必须分成两步：先创建文件夹、再修改权限
			os.Mkdir(cfg.Log.FileFolder, 0755) //0755也可以os.ModePerm
			os.Chmod(cfg.Log.FileFolder, 0755)
		}
	}

	// 根据配置文件配置的日志路径和文件名
	if cfg.Log.FileName == "" {
		panic("log.fileName is not correct in config yaml")
	}

	folderPath := filepath.Join(cfg.Log.FileFolder, cfg.Log.FileName)
	//_, err := os.Create(folderPath)
	_, err := os.OpenFile(folderPath, os.O_APPEND|os.O_CREATE, 0666) // os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)会覆盖或重新创建
	if err != nil {
		return nil
	}
	// 新增日志切割和归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   folderPath,         // 日志文件的位置
		MaxSize:    cfg.Log.MaxSize,    // 在进行切割之前，日志文件的最大大小
		MaxBackups: cfg.Log.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     cfg.Log.MaxAge,     // 保留旧文件的最大天数
		Compress:   cfg.Log.Compress,   // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GetLogInstance 获取log日志对象
func GetLogInstance() *zap.SugaredLogger {
	if sugarLogger != nil {
		return sugarLogger
	}
	return nil
}
