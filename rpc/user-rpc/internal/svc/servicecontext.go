package svc

import (
	"frozen-go-project/rpc/user-rpc/internal/config"
	model "frozen-go-project/rpc/user-rpc/internal/model/mysql"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c config.Config
	UserModel *model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
