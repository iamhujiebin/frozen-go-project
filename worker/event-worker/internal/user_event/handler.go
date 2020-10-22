package user_event

import (
	"github.com/Shopify/sarama"
	"github.com/tal-tech/go-zero/core/logx"
)

func HandleUserEvent(message *sarama.ConsumerMessage) bool {
	logx.Infof("UserEvent Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
	return true
}
