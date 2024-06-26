// Code generated by goctl. DO NOT EDIT.
// Source: foodiLBS.proto

package server

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"
	"github.com/foodi-org/foodi-lbs-server/internal/logic/delivery"
	"github.com/foodi-org/foodi-lbs-server/internal/svc"
)

type DeliveryServer struct {
	svcCtx *svc.ServiceContext
	foodi_lbs_server.UnimplementedDeliveryServer
}

func NewDeliveryServer(svcCtx *svc.ServiceContext) *DeliveryServer {
	return &DeliveryServer{
		svcCtx: svcCtx,
	}
}

func (s *DeliveryServer) Demo(ctx context.Context, in *foodi_lbs_server.DeliverDemoRequest) (*foodi_lbs_server.DeliverDemoReply, error) {
	l := deliverylogic.NewDemoLogic(ctx, s.svcCtx)
	return l.Demo(in)
}
