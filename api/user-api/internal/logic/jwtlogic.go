package logic

import (
	"context"
	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type JwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) JwtLogic {
	return JwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JwtLogic) Jwt(req types.JwtTokenRequest) (*types.JwtTokenResponse, error) {
	var accessExpire = l.svcCtx.Config.JwtAuth.AccessExpire

	now := time.Now().Unix()
	accessToken, err := l.GenToken(req.UserId, now, l.svcCtx.Config.JwtAuth.AccessSecret, nil, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.JwtTokenResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *JwtLogic) GenToken(userId, iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
