package commonconfig

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	commonConfigFieldNames          = FieldNames(&CommonConfig{})
	commonConfigRows                = strings.Join(commonConfigFieldNames, ",")
	commonConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(commonConfigFieldNames, "id", "create_time", "update_time"), ",")
	commonConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(commonConfigFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"

	cacheCommonConfigIdPrefix = "cache#CommonConfig#id#"
)

const dbTag = "db"

func FieldNames(in interface{}) []string {
	out := make([]string, 0)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(dbTag); tagv != "" {
			if tagv == "key" {
				out = append(out, fmt.Sprintf("`%s`", tagv))
			} else {
				out = append(out, tagv)
			}
		} else {
			out = append(out, fi.Name)
		}
	}
	return out
}

type (
	CommonConfigModel struct {
		db *sql.DB
		sqlc.CachedConn
		table string
	}

	CommonConfig struct {
		Id    int64  `db:"id"`
		Key   string `db:"key"`
		Value string `db:"value"`
	}
)

func NewCommonConfigModel(conn sqlx.SqlConn, db *sql.DB, c cache.CacheConf) *CommonConfigModel {
	return &CommonConfigModel{
		db:         db,
		CachedConn: sqlc.NewConn(conn, c, cache.WithExpiry(time.Minute)),
		table:      "common_config",
	}
}

func (m *CommonConfigModel) Insert(data CommonConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, commonConfigRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Key, data.Value)

	return ret, err
}

func (m *CommonConfigModel) FindOne(id int64) (*CommonConfig, error) {
	commonConfigIdKey := fmt.Sprintf("%s%v", cacheCommonConfigIdPrefix, id)
	var resp CommonConfig
	err := m.QueryRow(&resp, commonConfigIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = ? limit 1", commonConfigRows, m.table)
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

func (m *CommonConfigModel) Update(data CommonConfig) error {
	commonConfigIdKey := fmt.Sprintf("%s%v", cacheCommonConfigIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = ?", m.table, commonConfigRowsWithPlaceHolder)
		return conn.Exec(query, data.Key, data.Value, data.Id)
	}, commonConfigIdKey)
	return err
}

func (m *CommonConfigModel) Delete(id int64) error {

	commonConfigIdKey := fmt.Sprintf("%s%v", cacheCommonConfigIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = ?", m.table)
		return conn.Exec(query, id)
	}, commonConfigIdKey)
	return err
}

func (m *CommonConfigModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCommonConfigIdPrefix, primary)
}

func (m *CommonConfigModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", commonConfigRows, m.table)
	return conn.QueryRow(v, query, primary)
}
