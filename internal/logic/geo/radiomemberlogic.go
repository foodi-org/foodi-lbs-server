package geologic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"

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

// RadioMember
// @Description: 根据成员位置寻找半径范围内的坐标,按照距离从近到远排序。
// @param key redis geo set key
// @param member redis geo set member
// @param radius 搜索半径
// @param uint 搜索单位:m|km
// @param storeKey string 搜索结果存储的key
// @return []redis.GeoLocation 如: [GeoRadiusByMember] {X广场 104.07283455133438 30.663422957895442 0 4026137797693549}
// @return error
func (l *RadioMemberLogic) RadioMember(in *foodi_lbs_server.RadioMemberRequest) (*foodi_lbs_server.RadioOrderReply, error) {
	var (
		geoloc   []redis.GeoLocation
		err      error
		location = make([]*foodi_lbs_server.GeoLocation, 0)
	)
	if geoloc, err = l.svcCtx.Redis.GeoRadiusByMember(
		in.GetKey(),
		in.GetMember(),
		&redis.GeoRadiusQuery{
			Radius:      in.GetRadio(),
			Unit:        in.GetUnit().String(),
			WithCoord:   in.GetOption().GetWithCoord(),
			WithDist:    in.GetOption().GetWithDist(),
			WithGeoHash: in.GetOption().GetWithGeoHash(),
			Sort:        in.GetSort().String(),
			Count:       int(in.GetCount()),
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
