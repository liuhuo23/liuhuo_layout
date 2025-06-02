package service

import (
	"context"

	v1 "liuhuo23/liuos/api/health/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthSrv struct {
	v1.HealthSvcHTTPClientImpl
}

func NewHealthSrv() *HealthSrv {
	return &HealthSrv{}
}

func (h *HealthSrv) CheckHealth(context.Context, *emptypb.Empty) (*v1.HelloReply, error) {
	return &v1.HelloReply{
		Ok: "ok",
	}, nil
}
