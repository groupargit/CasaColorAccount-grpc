package logic

import (
	"context"

	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/svc"
	"github.com/groupargit/casacoloraccount-grpc/rpc/types/account/v1alpha1"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *v1alpha1.AccountLoginRequest) (*v1alpha1.AccountLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &v1alpha1.AccountLoginResponse{}, nil
}
