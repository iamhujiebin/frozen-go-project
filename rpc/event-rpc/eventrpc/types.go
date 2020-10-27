// Code generated by goctl. DO NOT EDIT!
// Source: event_rpc.proto

package eventrpc

import "errors"

var errJsonConvert = errors.New("json convert error")

type (
	CommonMessage struct {
		Topic       string `json:"topic,omitempty"`
		EventTimeMs int64  `json:"event_time_ms,omitempty"`
		Message     []byte `json:"message,omitempty"`
	}

	CommonResponse struct {
		Partition int32 `json:"partition,omitempty"`
		Offset    int64 `json:"offset,omitempty"`
	}

	UserAction struct {
		UserAction string `json:"user_action,omitempty"`
		UserId     int64  `json:"user_id,omitempty"`
	}

	UserActionReq struct {
		Common     *CommonMessage `json:"common,omitempty"`
		UserAction *UserAction    `json:"userAction,omitempty"`
	}
)
