package greetlogic

import (
	"context"

	"github.com/mikokutou1/go-zero-m/tools/goctl/example/rpc/hi/internal/svc"
	"github.com/mikokutou1/go-zero-m/tools/goctl/example/rpc/hi/pb/hi"

	"github.com/mikokutou1/go-zero-m/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *hi.HelloReq) (*hi.HelloResp, error) {
	// todo: add your logic here and delete this line

	return &hi.HelloResp{}, nil
}
