package geologic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeoPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeoPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeoPositionLogic {
	return &GeoPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeoPositionLogic) GeoPosition(in *foodi_lbs_server.PositionRequest) (*foodi_lbs_server.PositionReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_lbs_server.PositionReply{}, nil
}
