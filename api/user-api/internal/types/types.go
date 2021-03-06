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

type AnchorRecommendRequest struct {
	CommonParams
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}

type CommonParams struct {
	UserId      int64  `form:"__user_id,optional"`
	GuestId     string `form:"__guest_id,optional"`
	Country     string `form:"__country,optional"`
	PkgName     string `form:"__pname,optional"`
	PkgChannel  string `form:"__pch,optional"`
	UserChannel string `form:"__uch,optional"`
	Platform    string `form:"__platform,optional"`
	Version     string `form:"__v,optional"`
}

type CommonResponse struct {
	Code    int         `json:"code"`
	Body    interface{} `json:"body,omitempty"`
	Message string      `json:"message,omitempty"`
}

type DispatcherRequest struct {
	CommonParams
	Id  string `form:"id"`
	Pkg string `path:"pkg"`
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
	CommonParams
	AndroidId  string `form:"android_id,optional"`
	Imei       string `form:"imei,optional"`
	CampaignId string `form:"campaign_id,optional"`
}

type GuestLoginRequest struct {
	CommonParams
}

type HelloRequest struct {
}

type JwtTokenRequest struct {
	UserId int64 `path:"userId"`
}

type JwtTokenResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"` // 建议客户端刷新token的绝对时间
}

type UserActionRequest struct {
	CommonParams
	UserAction string `form:"user_action"`
}
