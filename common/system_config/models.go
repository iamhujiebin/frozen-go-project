package system_config

import "frozen-go-project/common/enum"

type CoinPromotionActionPoint map[enum.UserAction]int //金币促销行为分
type VipPromotionActionPoint map[enum.UserAction]int  //vip促销行为分
type GpPayChannelActionPoint map[enum.UserAction]int  //googlePay开启行为分
