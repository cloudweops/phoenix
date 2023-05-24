package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/CloudWeOps/phoenix/app"
	"github.com/CloudWeOps/phoenix/conf"
	"github.com/CloudWeOps/phoenix/logger"
	"github.com/cloudwego/hertz/pkg/route"
)

// HertzHTTPService hertz 服务
type HertzHTTPService struct {
	r      route.IRouter
	l      logger.Logger
	c      *conf.Config
	server *http.Server
}

func (s *HertzHTTPService) PathPrefix() string {
	return fmt.Sprintf("/%s/api", s.c.App.Name)
}

// Start 启动服务
func (s *HertzHTTPService) Start() error {
	// 装置子服务路由
	app.LoadHertzApp(s.PathPrefix(), s.r)

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
func (s *HertzHTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
