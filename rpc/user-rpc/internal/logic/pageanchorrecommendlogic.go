package logic

import (
	"context"
	"frozen-go-project/common/enum"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"

	"frozen-go-project/rpc/user-rpc/internal/svc"
	user_rpc "frozen-go-project/rpc/user-rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type PageAnchorRecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageAnchorRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageAnchorRecommendLogic {
	return &PageAnchorRecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PageAnchorRecommendLogic) PageAnchorRecommend(in *user_rpc.PageAnchorRecommendReq) (*user_rpc.PageAnchorRecommendRes, error) {
	where := bson.M{
		"user_role": enum.UserRoleEnum.Anchor,
	}
	res, err := l.svcCtx.UserMongoModel.PageUsers(where, in.Skip, in.Limit)
	if err != nil {
		return nil, err
	}
	var pbUserInfos []*user_rpc.UserInfo
	for k := range res {
		pbUser := new(user_rpc.UserInfo)
		_ = copier.Copy(pbUser, res[k])
		pbUserInfos = append(pbUserInfos, pbUser)
	}
	return &user_rpc.PageAnchorRecommendRes{
		Anchors: pbUserInfos,
	}, nil
}
