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
	guestsFieldNames          = builderx.FieldNames(&Guests{})
	guestsRows                = strings.Join(guestsFieldNames, ",")
	guestsRowsExpectAutoSet   = strings.Join(stringx.Remove(guestsFieldNames, "id", "create_time", "update_time"), ",")
	guestsRowsWithPlaceHolder = strings.Join(stringx.Remove(guestsFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"

	cacheGuestsIdPrefix = "cache#Guests#id#"
)

type (
	GuestsModel struct {
		sqlc.CachedConn
		table string
	}

	Guests struct {
		Id         int64     `db:"id"`          // 自增id
		GuestId    string    `db:"guest_id"`    // 游客id
		Platform   string    `db:"platform"`    // 平台：ios/android/web
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewGuestsModel(conn sqlx.SqlConn, c cache.CacheConf, table string) *GuestsModel {
	return &GuestsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}

func (m *GuestsModel) Insert(data Guests) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + guestsRowsExpectAutoSet + `) values (?, ?)`
	return m.ExecNoCache(query, data.GuestId, data.Platform)
}

func (m *GuestsModel) FindOne(id int64) (*Guests, error) {
	guestsIdKey := fmt.Sprintf("%s%v", cacheGuestsIdPrefix, id)
	var resp Guests
	err := m.QueryRow(&resp, guestsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + guestsRows + ` from ` + m.table + ` where id = ? limit 1`
		return conn.QueryRow(v, query, id)
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

func (m *GuestsModel) Update(data Guests) error {
	guestsIdKey := fmt.Sprintf("%s%v", cacheGuestsIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + m.table + ` set ` + guestsRowsWithPlaceHolder + ` where id = ?`
		return conn.Exec(query, data.GuestId, data.Platform, data.Id)
	}, guestsIdKey)
	return err
}

func (m *GuestsModel) Delete(id int64) error {

	guestsIdKey := fmt.Sprintf("%s%v", cacheGuestsIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `delete from ` + m.table + ` where id = ?`
		return conn.Exec(query, id)
	}, guestsIdKey)
	return err
}
