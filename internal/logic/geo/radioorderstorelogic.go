package geologic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RadioOrderStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRadioOrderStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RadioOrderStoreLogic {
	return &RadioOrderStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RadioOrderStore
// @Description: 返回搜索半径范围内成员列表并写入storeKey，按照距离从近到远排序。存储storeKey组成规则: 用户id-商家类型
// @param key redis geo set key
// @param longitude 当前坐标:经度
// @param latitude 当前坐标:纬度
// @param radius 搜索半径
// @param uint 搜索单位:m|km|ft|mi
// @param storeKey 存储key
// @return []redis.GeoLocation
// @return error
func (l *RadioOrderStoreLogic) RadioOrderStore(in *foodi_lbs_server.RadioOrderStoreRequest) (*foodi_lbs_server.RadioOrderReply, error) {
	var (
		geoloc   []redis.GeoLocation
		err      error
		location = make([]*foodi_lbs_server.GeoLocation, 0)
	)
	if geoloc, err = l.svcCtx.Redis.GeoRadius(
		in.GetKey(), in.GetLongitude(),
		in.GetLatitude(),
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
		l.Logger.Error(fmt.Sprintf("RadioOrderStore set storeKey Expire time error, storeKey=%s, error=%v", in.GetStoreKey(), err))
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
