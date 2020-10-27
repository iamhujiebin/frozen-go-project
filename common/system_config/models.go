package system_config

import "frozen-go-project/common/enum"

type CoinPromotionActionPoint map[enum.UserAction]int64 //金币促销行为分
type VipPromotionActionPoint map[enum.UserAction]int64  //vip促销行为分
type GpPayChannelActionPoint map[enum.UserAction]int64  //googlePay开启行为分
