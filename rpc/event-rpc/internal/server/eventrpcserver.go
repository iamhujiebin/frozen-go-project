// Code generated by goctl. DO NOT EDIT!
// Source: event_rpc.proto

package server

import (
	"context"

	"frozen-go-project/rpc/event-rpc/internal/logic"
	"frozen-go-project/rpc/event-rpc/internal/svc"
	event_rpc "frozen-go-project/rpc/event-rpc/pb"
)

type EventRpcServer struct {
	svcCtx *svc.ServiceContext
}

func NewEventRpcServer(svcCtx *svc.ServiceContext) *EventRpcServer {
	return &EventRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *EventRpcServer) UserAction(ctx context.Context, in *event_rpc.UserActionReq) (*event_rpc.CommonResponse, error) {
	l := logic.NewUserActionLogic(ctx, s.svcCtx)
	return l.UserAction(in)
}