package user_event

import (
	"context"
	"encoding/json"
	"frozen-go-project/common/mq_msg"
	"frozen-go-project/common/system_config"
	"frozen-go-project/rpc/base-rpc/baserpc"
	"frozen-go-project/rpc/user-rpc/userrpc"
	"frozen-go-project/worker/event-worker/internal/svc"
	"github.com/Shopify/sarama"
	"github.com/tal-tech/go-zero/core/logx"
)

func HandleUserEvent(message *sarama.ConsumerMessage) bool {
	logx.Infof("UserEvent Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
	var userAction mq_msg.UserActionMsg
	err := json.Unmarshal(message.Value, &userAction)
	if err != nil {
		logx.Errorf("user action message json fail:%s", err.Error())
		return false
	}
	coinPromotionActionPoint, vipPromotionActionPoint, gpPayChannelActionPoint, success := loadActionPointConfig()
	if !success {
		return true
	}
	coinPoint := coinPromotionActionPoint[userAction.UserAction]
	vipPoint := vipPromotionActionPoint[userAction.UserAction]
	gpPoint := gpPayChannelActionPoint[userAction.UserAction]
	if coinPoint > 0 || vipPoint > 0 || gpPoint > 0 {
		_, err = svc.ServiceCtx.UserRpc.AddActionPoint(context.Background(), &userrpc.AddActionPointReq{
			UserId:             userAction.UserId,
			PayChannelPoint:    gpPoint,
			VipPromotionPoint:  vipPoint,
			CoinPromotionPoint: coinPoint,
		})
		if err != nil {
			logx.Errorf("add user point fail:%s", err.Error())
			return false
		}
	} else {
		logx.Errorf("not a single point to add")
	}
	return true
}

func loadActionPointConfig() (coinPromotionActionPoint system_config.CoinPromotionActionPoint, vipPromotionActionPoint system_config.VipPromotionActionPoint,
	gpPayChannelActionPoint system_config.GpPayChannelActionPoint, success bool) {
	res, err := svc.ServiceCtx.BaseRpc.GetSystemConfigs(context.Background(), &baserpc.GetSystemConfigReq{
		Section: string(system_config.ConfigSectionEnum.ActionPoint),
	})
	if err != nil {
		logx.Errorf("rpc get user point config fail:%s", err.Error())
		success = false
		return
	}
	if res == nil || len(res.Items) <= 0 {
		success = false
		return
	}
	for _, v := range res.Items {
		if v.Key == string(system_config.ConfigKeyEnum.GPPayChannelActionPoint) {
			err := json.Unmarshal([]byte(v.Value), &gpPayChannelActionPoint)
			if err != nil {
				logx.Errorf("json fail:%s", err.Error())
			}
		}
		if v.Key == string(system_config.ConfigKeyEnum.VipPromotionActionPoint) {
			err := json.Unmarshal([]byte(v.Value), &vipPromotionActionPoint)
			if err != nil {
				logx.Errorf("json fail:%s", err.Error())
			}
		}
		if v.Key == string(system_config.ConfigKeyEnum.CoinPromotionActionPoint) {
			err := json.Unmarshal([]byte(v.Value), &coinPromotionActionPoint)
			if err != nil {
				logx.Errorf("json fail:%s", err.Error())
			}
		}
	}
	success = true
	return
}
