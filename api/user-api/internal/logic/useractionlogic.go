package logic

import (
	"context"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/common/mq_msg"
	"frozen-go-project/common/timex"
	"frozen-go-project/rpc/event-rpc/eventrpc"

	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserActionLogic {
	return UserActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserActionLogic) UserAction(req types.UserActionRequest) (*types.CommonResponse, error) {
	res, err := l.svcCtx.EventRpc.UserAction(l.ctx, &eventrpc.UserActionReq{
		Common: &eventrpc.CommonMessage{
			Topic:       mq_msg.UserActionTopic,
			EventTimeMs: timex.NowMs(),
			Message:     nil,
		},
		UserAction: &eventrpc.UserAction{
			UserAction: req.UserAction,
			UserId:     req.UserId,
		},
	})
	if err != nil {
		return &types.CommonResponse{
			Code:    resp_codes.RpcError,
			Message: err.Error(),
		}, nil
	}
	return &types.CommonResponse{
		Body: res,
	}, nil
}
