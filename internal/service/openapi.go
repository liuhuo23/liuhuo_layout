package service

import (
	"liuhuo23/liuos/internal/conf"
	"os"

	"github.com/go-kratos/kratos/v2/errors"

	"github.com/go-kratos/kratos/v2/transport/http"
)

type OpenapiService struct {
	server *conf.Server
}

func NewOpenapiService(server *conf.Server) *OpenapiService {
	return &OpenapiService{server: server}
}

func (s *OpenapiService) DownloadFile(ctx http.Context) error {
	openapi_path := s.server.OpenapiPath
	if openapi_path == "" {
		return errors.New(404, "", "file not found")
	}
	data, err := os.ReadFile(openapi_path)
	if err != nil {
		return errors.New(500, "", "file read error")
	}
	ctx.Response().Header().Set("Content-Type", "application/yaml")
	if _, err := ctx.Response().Write(data); err != nil {
		return errors.New(500, "", "file write error")
	}
	return nil
}
