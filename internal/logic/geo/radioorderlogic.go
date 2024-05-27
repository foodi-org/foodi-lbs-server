package geologic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RadioOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRadioOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RadioOrderLogic {
	return &RadioOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RadioOrderLogic) RadioOrder(in *foodi_lbs_server.RadioOrderRequest) (*foodi_lbs_server.RadioOrderReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_lbs_server.RadioOrderReply{}, nil
}
