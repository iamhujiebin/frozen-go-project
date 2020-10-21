package db_errors

import (
	"errors"
	"strings"
)

var DBNilRes = errors.New("db nil res")

func IllegalParams(str ...string) error {
	return errors.New("illegal params:" + strings.Join(str, ","))
}
