package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	// Logger 日志结构体，可以输出 结构化日志
	Logger *zap.Logger
	// SugarLogger 日志结构体，可以输出 结构化日志、非结构化日志
	SugarLogger *zap.SugaredLogger
)

// InitLog 初始化日志结构体 Logger 与 SugarLogger
func InitLog(logHome string, logLevel zapcore.Level, outToFile bool) {
	config := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "Logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
		},
	}

	var cores []zapcore.Core
	if outToFile {
		cores = newFileWriter(config, logHome, logLevel)
	} else {
		cores = newConsoleWriter(config, logLevel)
	}

	Logger = zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddStacktrace(zap.WarnLevel),
	)
	SugarLogger = Logger.Sugar()
}

// LoggerFlush 刷新 logger 缓冲区
func LoggerFlush() {
	err := Logger.Sync()
	if err != nil {
		// 记录到系统输出流
		log.Println("flush zap logger error : ", err.Error())
	}
}

func newConsoleWriter(config zapcore.EncoderConfig, logLevel zapcore.Level) []zapcore.Core {
	return []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			logLevel,
		),
	}
}

func newFileWriter(config zapcore.EncoderConfig, logHome string, logLevel zapcore.Level) []zapcore.Core {
	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})

	//自定义日志级别：自定义Warn级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	infoWriter := getLogFileWriter(logHome, "console.log")
	errWriter := getLogFileWriter(logHome, "error.log")

	// 设置gin框架的日志写入
	gin.DefaultWriter = infoWriter
	gin.DefaultErrorWriter = errWriter

	return []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config),
			zapcore.AddSync(infoWriter),
			infoLevel,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config),
			zapcore.AddSync(errWriter),
			warnLevel,
		),
	}
}

// getLogFileWriter 提供一个根据文件大小拆分日志文件的写入类
func getLogFileWriter(dir, filename string) io.Writer {
	return &lumberjack.Logger{
		Filename: filepath.Join(dir, filename),
		MaxSize: 1024, // 1G
		MaxAge:   15,
		Compress: true,
	}
}
