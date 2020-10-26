package enum

var UserActionEnum = UserActionEnumT{
	TriggerValue:            "trigger_value",
	EnterVipPurchasePage:    "enter_vip_purchase_page",
	CreateVipPurchaseOrder:  "create_vip_purchase_order",
	EnterCoinPurchasePage:   "enter_coin_purchase_page",
	CreateCoinPurchaseOrder: "create_coin_purchase_order",
	EnterAnchorProfile:      "enter_anchor_profile",
	SendMessage:             "send_message",
	InviteVoice:             "invite_voice",
	InviteVideo:             "invite_video",
	AnswerAutoVoice:         "answer_auto_voice",
	VoiceMatch:              "voice_match",
	VideoMatch:              "video_match",
	EnterChatPage:           "enter_chat_page",
	ClickAnchorList:         "click_anchor_list",
	PaySuccess:              "pay_success",
}

type UserAction string

type UserActionEnumT struct {
	TriggerValue            UserAction `desc:"下发促销活动阈值"`
	EnterVipPurchasePage    UserAction `desc:"进入会员购买界面"`
	CreateVipPurchaseOrder  UserAction `desc:"创建会员购买订单"`
	EnterCoinPurchasePage   UserAction `desc:"进入金币购买界面"`
	CreateCoinPurchaseOrder UserAction `desc:"创建金币购买订单"`
	EnterAnchorProfile      UserAction `desc:"进入主播个人页"`
	SendMessage             UserAction `desc:"发送消息"`
	InviteVoice             UserAction `desc:"发起语音通话"`
	InviteVideo             UserAction `desc:"发起视频通话"`
	AnswerAutoVoice         UserAction `desc:"接通自动语音"`
	VoiceMatch              UserAction `desc:"点击随机语音匹配"`
	VideoMatch              UserAction `desc:"点击随机视频匹配"`
	EnterChatPage           UserAction `desc:"进入聊天界面"`
	ClickAnchorList         UserAction `desc:"点击主播列表"`
	PaySuccess              UserAction `desc:"支付成功"`
}
