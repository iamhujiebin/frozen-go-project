package logic

import (
	"context"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/rpc/user-rpc/userrpc"

	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AnchorRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAnchorRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) AnchorRecommendLogic {
	return AnchorRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnchorRecommendLogic) AnchorRecommend(req types.AnchorRecommendRequest) (*types.CommonResponse, error) {
	res, err := l.svcCtx.UserRpc.PageAnchorRecommend(l.ctx, &userrpc.PageAnchorRecommendReq{
		UserId: req.UserId,
		Skip:   (req.Page - 1) * req.Size,
		Limit:  req.Size,
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
