package geologic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeoAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeoAddLogic {
	return &GeoAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeoAddLogic) GeoAdd(in *foodi_lbs_server.GeoAddRequest) (*foodi_lbs_server.GeoAddReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_lbs_server.GeoAddReply{}, nil
}
