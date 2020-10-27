package mq_msg

import "frozen-go-project/common/enum"

const UserActionTopic = "user-action"

type UserActionMsg struct {
	UserId      int64           `json:"user_id"`
	UserAction  enum.UserAction `json:"user_action"`
	EventTimeMs int64           `json:"event_time_ms"`
}
