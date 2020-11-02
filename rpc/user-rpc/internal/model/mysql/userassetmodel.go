package model

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userAssetFieldNames          = builderx.FieldNames(&UserAsset{})
	userAssetRows                = strings.Join(userAssetFieldNames, ",")
	userAssetRowsExpectAutoSet   = strings.Join(stringx.Remove(userAssetFieldNames, "id", "create_time", "update_time"), ",")
	userAssetRowsWithPlaceHolder = strings.Join(stringx.Remove(userAssetFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"

	cacheUserAssetUserIdPrefix = "cache#UserAsset#userId#"
	cacheUserAssetIdPrefix     = "cache#UserAsset#id#"
)

type (
	UserAssetModel struct {
		sqlc.CachedConn
		db    *gorm.DB
		table string
	}

	UserAsset struct {
		Id                    int64      `db:"id"`
		UserId                int64      `db:"user_id"`          // 用户id
		AvailableCoin         int64      `db:"available_coin"`   // 可用的金币数
		AccumulatedCoin       int64      `db:"accumulated_coin"` // 累计金币数
		FreeChatTimes         int64      `db:"free_chat_times"`  // 免费聊天次数
		FreeCallMinute        int64      `db:"free_call_minute"` // 免费通话时长
		Version               int64      `db:"version"`          // 数据版本号，乐观锁
		CreateTime            time.Time `db:"create_time"`
		UpdateTime            time.Time `db:"update_time"`
		ExtNum                int64      `db:"ext_num"`
		ExtStr                string     `db:"ext_str"`
		VipEffectEnd          time.Time `db:"vip_effect_end"`
		AvailableSilverCoin   int64      `db:"available_silver_coin"`   // 可用的银币数
		AccumulatedSilverCoin int64      `db:"accumulated_silver_coin"` // 累计银币数
	}
)

func (u *UserAsset) TableName() string {
	tbIndex := u.UserId % 2 //分表
	if tbIndex == 0 {
		tbIndex = 2
	}
	table := fmt.Sprintf("user_asset_%d", tbIndex)
	return table
}

func NewUserAssetModel(conn sqlx.SqlConn, db *gorm.DB, c cache.CacheConf, table string) *UserAssetModel {
	return &UserAssetModel{
		db:         db,
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}

func (m *UserAssetModel) TableName(userId int64) string {
	tbIndex := userId % 2 //分表
	if tbIndex == 0 {
		tbIndex = 2
	}
	table := fmt.Sprintf("%s_%d", m.table, tbIndex)
	return table
}

func (m *UserAssetModel) Insert(data UserAsset) (sql.Result, error) {
	userAssetUserIdKey := fmt.Sprintf("%s%v", cacheUserAssetUserIdPrefix, data.UserId)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `insert into ` + m.TableName(data.UserId) + ` (` + userAssetRowsExpectAutoSet + `) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		return conn.Exec(query, data.UserId, data.AvailableCoin, data.AccumulatedCoin, data.FreeChatTimes, data.FreeCallMinute, data.Version, data.ExtNum, data.ExtStr, data.VipEffectEnd, data.AvailableSilverCoin, data.AccumulatedSilverCoin)
	}, userAssetUserIdKey)
	return ret, err
}

func (m *UserAssetModel) Insert2(data UserAsset) (UserAsset, error) {
	res := m.db.Create(&data)
	if res.Error != nil {
		return data, res.Error
	}
	return data, nil
}

func (m *UserAssetModel) FindOne(id int64) (*UserAsset, error) {
	userAssetIdKey := fmt.Sprintf("%s%v", cacheUserAssetIdPrefix, id)
	var resp UserAsset
	err := m.QueryRow(&resp, userAssetIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + userAssetRows + ` from ` + m.TableName(id) + ` where id = ? limit 1`
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

func (m *UserAssetModel) FindOneByUserId(userId int64) (*UserAsset, error) {
	userAssetUserIdKey := fmt.Sprintf("%s%v", cacheUserAssetUserIdPrefix, userId)
	var resp UserAsset
	err := m.QueryRowIndex(&resp, userAssetUserIdKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := `select ` + userAssetRows + ` from ` + m.TableName(userId) + ` where user_id = ? limit 1`
		if err := conn.QueryRow(&resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UserAssetModel) Update(data UserAsset) (sql.Result, error) {
	userAssetIdKey := fmt.Sprintf("%s%v", cacheUserAssetIdPrefix, data.Id)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + m.table + ` set ` + userAssetRowsWithPlaceHolder + ` where id = ?`
		return conn.Exec(query, data.UserId, data.AvailableCoin, data.AccumulatedCoin, data.FreeChatTimes, data.FreeCallMinute, data.Version, data.ExtNum, data.ExtStr, data.VipEffectEnd, data.AvailableSilverCoin, data.AccumulatedSilverCoin, data.Id)
	}, userAssetIdKey)
	return ret, err
}

func (m *UserAssetModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	userAssetUserIdKey := fmt.Sprintf("%s%v", cacheUserAssetUserIdPrefix, data.UserId)
	userAssetIdKey := fmt.Sprintf("%s%v", cacheUserAssetIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `delete from ` + m.table + ` where id = ?`
		return conn.Exec(query, id)
	}, userAssetUserIdKey, userAssetIdKey)
	return err
}

func (m *UserAssetModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserAssetIdPrefix, primary)
}

func (m *UserAssetModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := `select ` + userAssetRows + ` from ` + m.table + ` where id = ? limit 1`
	return conn.QueryRow(v, query, primary)
}
