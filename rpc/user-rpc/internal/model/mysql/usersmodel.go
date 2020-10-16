package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	usersFieldNames          = builderx.FieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "user_id", "create_time", "update_time"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "user_id", "create_time", "update_time"), "=?,") + "=?"

	cacheUsersUserIdPrefix = "cache#Users#userId#"
)

type (
	UsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		UserId      int64     `db:"user_id"`      // 用户id
		AccessToken string    `db:"access_token"` // 用户校验token
		Avatar      string    `db:"avatar"`       // 用户头像
		CreateTime  time.Time `db:"create_time"`  // 创建时间
		UpdateTime  time.Time `db:"update_time"`
	}
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) *UsersModel {
	return &UsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      UsersModel{}.TableName(),
	}
}

func (UsersModel) TableName() string {
	return "users"
}

func (m *UsersModel) Insert(data Users) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + usersRowsExpectAutoSet + `) values (?, ?)`
	return m.ExecNoCache(query, data.AccessToken, data.Avatar)
}

func (m *UsersModel) FindOne(userId int64) (*Users, error) {
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, userId)
	var resp Users
	err := m.QueryRow(&resp, usersUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + usersRows + ` from ` + m.table + ` where user_id = ? limit 1`
		return conn.QueryRow(v, query, userId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UsersModel) Update(data Users) error {
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + m.table + ` set ` + usersRowsWithPlaceHolder + ` where user_id = ?`
		return conn.Exec(query, data.AccessToken, data.Avatar, data.UserId)
	}, usersUserIdKey)
	return err
}

func (m *UsersModel) Delete(userId int64) error {

	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, userId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `delete from ` + m.table + ` where user_id = ?`
		return conn.Exec(query, userId)
	}, usersUserIdKey)
	return err
}
