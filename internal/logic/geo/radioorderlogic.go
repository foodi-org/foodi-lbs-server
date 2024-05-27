package geologic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"

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

// RadioOrder
// @Description: 返回坐标位置半径范围内的坐标,按照距离从进到远排序.以二维数组形式返回附带geo相关内容
// @param longitude 当前坐标:经度
// @param latitude 当前坐标:纬度
// @param radius 搜索半径
// @param unit 搜索单位:m|km
// @return []redis.GeoLocation 如: [GeoRadiusByMember] {X广场 104.07283455133438 30.663422957895442 0 4026137797693549}
// @return error
func (l *RadioOrderLogic) RadioOrder(in *foodi_lbs_server.RadioOrderRequest) (*foodi_lbs_server.RadioOrderReply, error) {
	var (
		geoloc   []redis.GeoLocation
		err      error
		location = make([]*foodi_lbs_server.GeoLocation, 0)
	)
	if in.GetOption() == nil {
		return nil, errors.New("option of uint must")
	}
	if geoloc, err = l.svcCtx.Redis.GeoRadius(
		in.GetKey(),
		in.GetLongitude(),
		in.GetLatitude(),
		&redis.GeoRadiusQuery{
			Radius:      in.GetRadius(),
			Unit:        in.GetLengthType().String(),
			WithCoord:   in.GetOption().GetWithCoord(),
			WithDist:    in.GetOption().GetWithDist(),
			WithGeoHash: in.GetOption().GetWithGeoHash(),
			Sort:        in.GetOption().GetSort(),
		},
	); err != nil {
		return nil, err
	}
	for _, loc := range geoloc {
		location = append(location, &foodi_lbs_server.GeoLocation{
			Name:      loc.Name,
			Longitude: loc.Longitude,
			Latitude:  loc.Latitude,
			Dist:      loc.Dist,
			GeoHash:   loc.GeoHash,
		})
	}
	return &foodi_lbs_server.RadioOrderReply{Location: location}, nil
}
