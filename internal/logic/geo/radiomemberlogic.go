package geologic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RadioMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRadioMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RadioMemberLogic {
	return &RadioMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RadioMemberLogic) RadioMember(in *foodi_lbs_server.RadioMemberRequest) (*foodi_lbs_server.RadioOrderReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_lbs_server.RadioOrderReply{}, nil
}
