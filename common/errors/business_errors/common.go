package business_errors

import (
	"errors"
	"strings"
)

func EmptyParams(params ...string) error {
	return errors.New("empty params:" + strings.Join(params, ","))
}

var NoConfigErr = errors.New("no config")
var NoMqTopic = errors.New("no mq topic")
