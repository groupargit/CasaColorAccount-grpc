package logic

import (
	"context"
	"fmt"

	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLoggerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLoggerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLoggerLogic {
	return &CreateLoggerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLoggerLogic) CreateLogger(req []byte) error {
	fmt.Println("CreateLoggerLogic")
	fmt.Println("req===:", string(req))

	//push data in folder storagelog
	createLogger := l.svcCtx.DB.CreateLogger(l.ctx, req)
	if createLogger != nil {
		return createLogger
	}

	return nil
}
