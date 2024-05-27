package geologic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RadioOrderStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRadioOrderStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RadioOrderStoreLogic {
	return &RadioOrderStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RadioOrderStoreLogic) RadioOrderStore(in *foodi_lbs_server.RadioOrderStoreRequest) (*foodi_lbs_server.RadioOrderReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_lbs_server.RadioOrderReply{}, nil
}
