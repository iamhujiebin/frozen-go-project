package user_event

import (
	"context"
	"frozen-go-project/common/system_config"
	"frozen-go-project/rpc/base-rpc/baserpc"
	"frozen-go-project/worker/event-worker/internal/svc"
	"github.com/Shopify/sarama"
	"github.com/tal-tech/go-zero/core/logx"
)

func HandleUserEvent(message *sarama.ConsumerMessage) bool {
	logx.Infof("UserEvent Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
	res, err := svc.ServiceCtx.BaseRpc.GetSystemConfigs(context.Background(), &baserpc.GetSystemConfigReq{
		Section: string(system_config.ConfigSectionEnum.ActionPoint),
	})
	logx.Infof("test rpc:%v,%v", res, err)
	return true
}
