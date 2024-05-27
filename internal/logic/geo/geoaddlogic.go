package geologic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"

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
	if res, err := l.svcCtx.Redis.GeoAdd(in.Key, &redis.GeoLocation{
		Name:      in.GetName(),
		Longitude: in.GetLongitude(),
		Latitude:  in.GetLatitude(),
	}); err != nil {
		return nil, err
	} else {
		return &foodi_lbs_server.GeoAddReply{Idx: res}, nil
	}
}
