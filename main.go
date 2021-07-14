package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yuefei7746/go-web-example/utils"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	serverPort            = ":8080"
	serverReadTimeout     = 5 * time.Second
	serverWriteTimeout    = 5 * time.Second
	serverShutdownTimeout = 3 * time.Second
	serverMaxHeaderBytes  = 1 << 20

	ProductionMode  = "prod"
	DevelopmentMode = "dev"
)

func main() {
	initArgs()

	r := initRouter()
	listenAndServe(r)
}

func listenAndServe(router *gin.Engine) {
	srv := &http.Server{
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    serverReadTimeout,
		WriteTimeout:   serverWriteTimeout,
		MaxHeaderBytes: serverMaxHeaderBytes,
	}
	utils.Logger.Debug("Listening server")
	utils.Logger.Debug("web server config",
		zap.Duration("ReadTimeout", serverReadTimeout),
		zap.Duration("WriteTimeout", serverWriteTimeout),
		zap.Int("MaxHeaderBytes", serverMaxHeaderBytes),
	)

	// 在goroutine中初始化服务器，以便它不会阻止下面的正常关闭处理
	go func() {
		utils.SugarLogger.Infof("Serving on port %s", serverPort)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			utils.SugarLogger.Errorf("listen: %v", err)
		}
	}()

	// 等待中断信号正常关闭服务器
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Logger.Info("Shutting down server...")

	// 上下文用于通知服务器它有一定的时间来完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		// 被迫关闭
		utils.SugarLogger.Fatalf("Server forced to shutdown: %v", err)
	}

	utils.Logger.Info("Server exiting")
	utils.LoggerFlush()
}
