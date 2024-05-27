package geologic

import (
	"context"
	"errors"

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

/*
GeoDist
@Description: 计算两地相差距离
@param key: redis geo set key
@param in: in.member1 in.member2 位置名称
@return *foodiLBS.DistReply: DistReply.Dist 距离长度， DistReply.LengthType 单位
@return error
*/
func (l *GeoDistLogic) GeoDist(in *foodi_lbs_server.DistRequest) (*foodi_lbs_server.DistReply, error) {
	if in.GetLengthType() != foodi_lbs_server.LengthCorpus_KILOMETER && in.GetLengthType() != foodi_lbs_server.LengthCorpus_METER {
		return nil, errors.New("错误的单位")
	}
	if d, err := l.svcCtx.Redis.GeoDist(in.GetKey(), in.GetMember1(), in.GetMember2(), in.GetLengthType().String()); err != nil {
		return nil, err
	} else {
		return &foodi_lbs_server.DistReply{
			Dist:       d,
			LengthType: in.GetLengthType().String(),
		}, nil
	}
}
