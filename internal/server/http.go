package server

import (
	health_v1 "liuhuo23/liuos/api/health/v1"
	"liuhuo23/liuos/internal/conf"
	"liuhuo23/liuos/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	logger log.Logger,
	health *service.HealthSrv,
	openapi *service.OpenapiService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	health_v1.RegisterHealthSvcHTTPServer(srv, health)
	srv.Route("/").GET("/openapi.yaml", openapi.DownloadFile)
	return srv
}
