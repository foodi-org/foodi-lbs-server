package geologic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"

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

// RadioMemberStore
// @Description: 根据成员位置寻找半径范围内的坐标,按照距离从近到远排序。并将结果写入store便于分页返回。storeKey Expire time 20min
// @param key redis geo set key
// @param member redis geo set member
// @param radius 搜索半径
// @param unit 搜索单位:m|km
// @param storeKey string 搜索结果存储的key
// @return []redis.GeoLocation 如: [GeoRadiusByMember] {X广场 104.07283455133438 30.663422957895442 0 4026137797693549}
// @return error
func (l *RadioMemberStoreLogic) RadioMemberStore(in *foodi_lbs_server.RadioMemberStoreRequest) (*foodi_lbs_server.RadioOrderReply, error) {
	var (
		geoloc   []redis.GeoLocation
		err      error
		location = make([]*foodi_lbs_server.GeoLocation, 0)
	)
	if geoloc, err = l.svcCtx.Redis.GeoRadiusByMember(
		in.GetKey(),
		in.GetMember(),
		&redis.GeoRadiusQuery{
			Radius:    in.GetRadius(),        //radius表示范围距离
			Unit:      in.GetLengthType(),    //距离单位是 m|km|ft|mi
			Sort:      in.GetSort().String(), //默认结果是未排序的，传入ASC为从近到远排序，传入DESC为从远到近排序
			StoreDist: in.GetStoreKey(),      // 将搜索结果存入storeKey, Store不按照排序存入,StoreDist按照顺序存入
		}); err != nil {
		return nil, err
	}

	// 设置过期时间
	err = l.svcCtx.Redis.Expire(in.GetStoreKey(), 1200)
	if err != nil {
		l.Logger.Error(fmt.Sprintf("RadioMemberStore set storeKey Expire time error, storeKey=%s, error=%v", in.GetStoreKey(), err))
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
