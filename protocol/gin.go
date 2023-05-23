package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/CloudWeOps/phoenix/app"
	"github.com/CloudWeOps/phoenix/conf"
	"github.com/CloudWeOps/phoenix/http/middleware/cors"
	"github.com/CloudWeOps/phoenix/logger"
	"github.com/CloudWeOps/phoenix/logger/zap"
	"github.com/gin-gonic/gin"
)

// NewHTTPService 构建函数
func NewGinHTTPService() *GinHTTPService {

	r := gin.New()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HTTP.Addr(),
		Handler:           cors.AllowAll().Handler(r),
	}
	return &GinHTTPService{
		r:      r,
		server: server,
		l:      zap.L().Named("HTTP Service"),
		c:      conf.C(),
	}
}

// HTTPService http服务
type GinHTTPService struct {
	r      gin.IRouter
	l      logger.Logger
	c      *conf.Config
	server *http.Server
}

func (s *GinHTTPService) PathPrefix() string {
	return fmt.Sprintf("/%s/api", s.c.App.Name)
}

// Start 启动服务
func (s *GinHTTPService) Start() error {
	// 装置子服务路由
	app.LoadGinApp(s.PathPrefix(), s.r)

	// 启动 HTTP服务
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Stop 停止server
func (s *GinHTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
