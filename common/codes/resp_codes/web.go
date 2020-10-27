package resp_codes

var RpcError = 1001
var CheckAccessTokenError = 1002

var CheckAccessTokenFail = struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}{
	Code:    CheckAccessTokenError,
	Message: "check accessToken fail",
}
