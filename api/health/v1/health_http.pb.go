// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v4.25.1
// source: health/v1/health.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationHealthSvcCheckHealth = "/health.v1.healthSvc/CheckHealth"

type HealthSvcHTTPServer interface {
	CheckHealth(context.Context, *emptypb.Empty) (*HelloReply, error)
}

func RegisterHealthSvcHTTPServer(s *http.Server, srv HealthSvcHTTPServer) {
	r := s.Route("/")
	r.GET("/health", _HealthSvc_CheckHealth0_HTTP_Handler(srv))
}

func _HealthSvc_CheckHealth0_HTTP_Handler(srv HealthSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationHealthSvcCheckHealth)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckHealth(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type HealthSvcHTTPClient interface {
	CheckHealth(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type HealthSvcHTTPClientImpl struct {
	cc *http.Client
}

func NewHealthSvcHTTPClient(client *http.Client) HealthSvcHTTPClient {
	return &HealthSvcHTTPClientImpl{client}
}

func (c *HealthSvcHTTPClientImpl) CheckHealth(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/health"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationHealthSvcCheckHealth))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
