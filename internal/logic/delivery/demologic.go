package deliverylogic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DemoLogic) Demo(in *foodi_lbs_server.DeliverDemoRequest) (*foodi_lbs_server.DeliverDemoReply, error) {

	return &foodi_lbs_server.DeliverDemoReply{Regis: true}, nil
}
