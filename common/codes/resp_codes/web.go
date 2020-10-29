package resp_codes

var RpcError = 1001
var CheckAccessTokenError = 1002
var CheckSignError = 1003
var BanError = 1004

var CheckAccessTokenFail = struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}{
	Code:    CheckAccessTokenError,
	Message: "check accessToken fail",
}

var CheckSignFail = struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}{
	Code:    CheckSignError,
	Message: "check sign fail",
}

var BanFail = struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}{
	Code:    BanError,
	Message: "ban !",
}
