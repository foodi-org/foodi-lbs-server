package geologic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeoDistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeoDistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeoDistLogic {
	return &GeoDistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeoDistLogic) GeoDist(in *foodi_lbs_server.DistRequest) (*foodi_lbs_server.DistReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_lbs_server.DistReply{}, nil
}
