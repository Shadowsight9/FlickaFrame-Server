package logic

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/pb/user_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *user_service.FindByIdRequest) (*user_service.FindByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &user_service.FindByIdResponse{}, nil
}
