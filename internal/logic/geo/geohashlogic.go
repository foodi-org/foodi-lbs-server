package geologic

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeoHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeoHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeoHashLogic {
	return &GeoHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GeoHash
// @Description: 获取位置的geo hash string value
// @param key redis geo set key
// @param member 位置
// @return []string 如:[GeoHash] [wm6n2npe2u0]
// @return error
func (l *GeoHashLogic) GeoHash(in *foodi_lbs_server.HashRequest) (*foodi_lbs_server.HashReply, error) {
	res, err := l.svcCtx.Redis.GeoHash(in.GetKey(), in.GetMembers()...)
	if err != nil {
		return nil, err
	}

	return &foodi_lbs_server.HashReply{Hash: res}, nil
}
