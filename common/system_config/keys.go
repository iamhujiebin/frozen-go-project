package system_config

// 系统配置key枚举的值类型
type ConfigKey string

type configKeyEnumT struct {
	VipPromotionActionPoint  ConfigKey //会员促销用户积分配置（json格式），积分达标后可见会员促销
	CoinPromotionActionPoint ConfigKey //金币促销用户积分配置（json格式），积分达标后可见金币促销
	GPPayChannelActionPoint  ConfigKey //支付渠道用户积分配置（json格式），积分达标后可见所有支付渠道，主要针对GooglePlay用户
}

var ConfigKeyEnum = configKeyEnumT{
	VipPromotionActionPoint:  "vip_promotion_action_point",
	CoinPromotionActionPoint: "coin_promotion_action_point",
	GPPayChannelActionPoint:  "gp_pay_channel_action_point",
}
