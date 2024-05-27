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

func (l *GeoHashLogic) GeoHash(in *foodi_lbs_server.HashRequest) (*foodi_lbs_server.HashReply, error) {
	// todo: add your logic here and delete this line

	return &foodi_lbs_server.HashReply{}, nil
}
