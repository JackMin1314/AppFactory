package log

import (
	"AppFactory/pkg/config"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	klog "github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLog struct {
	sugarLogger *zap.SugaredLogger
	pool        *sync.Pool
}

var log = &ZapLog{
	pool: &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	},
}

var once sync.Once

// InitLogger 日志初始化
func InitLogger(cfgYml *config.ConfigYaml) {
	once.Do(func() {
		writeSyncer := getLogWriter(cfgYml)
		encoder := getEncoder()
		var core zapcore.Core
		if cfgYml.Application.Env == "dev" {
			core = zapcore.NewTee(
				zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel),
				zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 终端和日志都记录
			)
		} else {
			core = zapcore.NewTee(
				zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel),
			)
		}

		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 封装了日志对象，所有打印的时候需要跳一层
		defer logger.Sync()
		log.sugarLogger = logger.Sugar()
		log.sugarLogger.Info("first init logger success")
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
func GetLogInstance() *ZapLog {
	if log != nil && log.sugarLogger != nil {
		return log
	}
	return nil
}

// Print 封装下Print结构化的日志
func (sug *ZapLog) Print(kvpair ...interface{}) {
	if len(kvpair) == 0 {
		return
	}
	if len(kvpair)%2 != 0 {
		kvpair = append(kvpair, "")
	}
	buf := sug.pool.Get().(*bytes.Buffer)
	var logT klog.Level
	for i := 0; i < len(kvpair); i += 2 {
		fmt.Fprintf(buf, "%s=%v ", kvpair[i], kvpair[i+1])
		value, ok := kvpair[i+1].(klog.Level)
		if ok {
			logT = value
		}
	}
	switch logT {
	case klog.LevelDebug:
		sug.sugarLogger.Debug(buf.String())
	case klog.LevelInfo:
		sug.sugarLogger.Info(buf.String())
	case klog.LevelWarn:
		sug.sugarLogger.Warn(buf.String())
	case klog.LevelError:
		sug.sugarLogger.Error(buf.String())
	default:
		sug.sugarLogger.Info(buf.String())
	}
	buf.Reset()
	sug.pool.Put(buf)
}

func (sug *ZapLog) Info(args ...interface{}) {
	sug.sugarLogger.Info(args)
}

func (sug *ZapLog) Infof(template string, args ...interface{}) {
	sug.sugarLogger.Infof(template, args...)
}

func (sug *ZapLog) Warn(args ...interface{}) {
	sug.sugarLogger.Warn(args)
}

func (sug *ZapLog) Warnf(template string, args ...interface{}) {
	sug.sugarLogger.Warnf(template, args...)
}

func (sug *ZapLog) Error(args ...interface{}) {
	sug.sugarLogger.Error(args)
}

func (sug *ZapLog) Errorf(template string, args ...interface{}) {
	sug.sugarLogger.Errorf(template, args...)
}

func (sug *ZapLog) Errorw(msg string, keysAndValues ...interface{}) {
	sug.sugarLogger.Errorw(msg, keysAndValues...)
}

func (sug *ZapLog) Debug(args ...interface{}) {
	sug.sugarLogger.Debug(args)
}

func (sug *ZapLog) Debugf(template string, args ...interface{}) {
	sug.sugarLogger.Debugf(template, args...)
}

func (sug *ZapLog) Fatal(args ...interface{}) {
	sug.sugarLogger.Fatal(args...)
}

func (sug *ZapLog) Fatalf(template string, args ...interface{}) {
	sug.sugarLogger.Fatalf(template, args...)
}
