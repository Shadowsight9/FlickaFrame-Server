// Code generated by goctl. DO NOT EDIT.
// Source: oss.proto

package oss

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/pb/oss_service"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateUpTokenRequest  = oss_service.CreateUpTokenRequest
	CreateUpTokenResponse = oss_service.CreateUpTokenResponse

	Oss interface {
		CreatUpToken(ctx context.Context, in *CreateUpTokenRequest, opts ...grpc.CallOption) (*CreateUpTokenResponse, error)
	}

	defaultOss struct {
		cli zrpc.Client
	}
)

func NewOss(cli zrpc.Client) Oss {
	return &defaultOss{
		cli: cli,
	}
}

func (m *defaultOss) CreatUpToken(ctx context.Context, in *CreateUpTokenRequest, opts ...grpc.CallOption) (*CreateUpTokenResponse, error) {
	client := oss_service.NewOssClient(m.cli.Conn())
	return client.CreatUpToken(ctx, in, opts...)
}
