// Code generated by goctl. DO NOT EDIT.
package types

type AddUserRequest struct {
	Avatar string `json:"avatar"`
}

type AddUserResponse struct {
	UserId         string `json:"user_id"`
	Avatar         string `json:"avatar"`
	AccessToken    string `json:"access_token"`
	CreateTimeUnix int64  `json:"create_time_unix"`
}

type CommonResponse struct {
	Code    int         `json:"code"`
	Body    interface{} `json:"body"`
	Message string      `json:"message"`
}

type GetUserRequest struct {
	UserId string `json:"userId"`
}

type GetUserResponse struct {
	UserId         string `json:"userId"`
	AccessToken    string `json:"accessToken"`
	Avatar         string `json:"avatar"`
	CreateTimeUnix int64  `json:"createTimeUnix"`
}

type GuestInitRequest struct {
	UserId      int    `form:"user_id,optional"`
	GuestId     string `form:"guest_id,optional"`
	GuestName   string `form:"guestname,optional"`
	Platform    string `form:"platform"`
	AndroidId   string `form:"android_id,optional"`
	AppVersion  string `form:"app_version"`
	Country     string `form:"country,optional"`
	Imei        string `form:"imei,optional"`
	Channel     string `form:"channel,optional"`
	UserChannel string `form:"user_channel,optional"`
	CampaignId  string `form:"campaign_id,optional"`
}

type JwtTokenRequest struct {
}

type JwtTokenResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"` // 建议客户端刷新token的绝对时间
}
