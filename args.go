package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/yuefei7746/go-web-example/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initArgs() {
	appMode := flag.String("appMode", DevelopmentMode, "application running mode, must in [prod,dev]")
	logHome := flag.String("logHome", "logs", "application log home")
	flag.Parse()

	var logLevel zapcore.Level
	var outToFile bool
	switch *appMode {
	case ProductionMode:
		gin.SetMode(gin.ReleaseMode)
		logLevel = zapcore.InfoLevel
		outToFile = true
		break
	case DevelopmentMode:
		gin.SetMode(gin.DebugMode)
		logLevel = zapcore.DebugLevel
		break
	default:
		panic("appMode must in [prod|dev]")
	}

	utils.InitLog(*logHome, logLevel, outToFile)

	utils.Logger.Info("application start ~~~")
	utils.Logger.Debug("application config",
		zap.String("appMode", *appMode),
		zap.String("logHome", *logHome),
		zap.Bool("outToFile", outToFile),
		zap.Int8("logLevel", int8(logLevel)),
	)
}
