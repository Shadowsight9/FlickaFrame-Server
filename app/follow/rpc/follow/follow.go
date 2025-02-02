// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package follow

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FansItem           = pb.FansItem
	FansListRequest    = pb.FansListRequest
	FansListResponse   = pb.FansListResponse
	FollowItem         = pb.FollowItem
	FollowListRequest  = pb.FollowListRequest
	FollowListResponse = pb.FollowListResponse
	FollowRequest      = pb.FollowRequest
	FollowResponse     = pb.FollowResponse
	IsFollowReq        = pb.IsFollowReq
	IsFollowResp       = pb.IsFollowResp
	UnFollowRequest    = pb.UnFollowRequest
	UnFollowResponse   = pb.UnFollowResponse

	Follow interface {
		Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
		UnFollow(ctx context.Context, in *UnFollowRequest, opts ...grpc.CallOption) (*UnFollowResponse, error)
		FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error)
		FansList(ctx context.Context, in *FansListRequest, opts ...grpc.CallOption) (*FansListResponse, error)
		IsFollow(ctx context.Context, in *IsFollowReq, opts ...grpc.CallOption) (*IsFollowResp, error)
	}

	defaultFollow struct {
		cli zrpc.Client
	}
)

func NewFollow(cli zrpc.Client) Follow {
	return &defaultFollow{
		cli: cli,
	}
}

func (m *defaultFollow) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

func (m *defaultFollow) UnFollow(ctx context.Context, in *UnFollowRequest, opts ...grpc.CallOption) (*UnFollowResponse, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.UnFollow(ctx, in, opts...)
}

func (m *defaultFollow) FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.FollowList(ctx, in, opts...)
}

func (m *defaultFollow) FansList(ctx context.Context, in *FansListRequest, opts ...grpc.CallOption) (*FansListResponse, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.FansList(ctx, in, opts...)
}

func (m *defaultFollow) IsFollow(ctx context.Context, in *IsFollowReq, opts ...grpc.CallOption) (*IsFollowResp, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.IsFollow(ctx, in, opts...)
}
