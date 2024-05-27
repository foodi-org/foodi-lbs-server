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

/*
GeoPosition
@Description: 获取位置经纬度
@param key redis geo set key
@param members 位置名称
@return []*redis.GeoPos 如:[GeoPos] Longitude :  104.07283455133438 Latitude :  30.663422957895442
@return error
*/
func (l *GeoPositionLogic) GeoPosition(in *foodi_lbs_server.PositionRequest) (*foodi_lbs_server.PositionReply, error) {
	pos, err := l.svcCtx.Redis.GeoPos(in.GetKey(), in.GetMembers()...)
	if err != nil {
		return nil, err
	}
	data := &foodi_lbs_server.PositionReply{Pos: make([]*foodi_lbs_server.Position, len(pos))}
	for k, v := range pos {
		data.Pos[k] = &foodi_lbs_server.Position{
			Longitude: v.Longitude,
			Latitude:  v.Latitude,
		}
	}

	return data, nil
}
