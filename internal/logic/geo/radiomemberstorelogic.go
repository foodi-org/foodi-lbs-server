package geologic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RadioMemberStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRadioMemberStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RadioMemberStoreLogic {
	return &RadioMemberStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RadioMemberStoreLogic) RadioMemberStore(in *foodi_lbs_server.RadioMemberStoreRequest) (*foodi_lbs_server.RadioOrderReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_lbs_server.RadioOrderReply{}, nil
}
