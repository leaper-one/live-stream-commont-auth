// Code generated by goctl. DO NOT EDIT.
package types

type ApplyNewVCodeRequest struct {
	Buid     int    `path:"buid"`
	Key      string `form:"key,optional"` // 不填 为 devloper 登录, 只有 经过弹幕认证的用户才能申请 vcode
	ClientID string `form:"client_id"`
}

type ApplyNewVCodeResponse struct {
	Vcode string `json:"vcode"`
}

type AddOneRequest struct {
	Vcode string `path:"vcode"`
	Buid  int    `json:"buid"`
}

type AddOneResponse struct {
	Status int `json:"status"`
}

type VerifyRequest struct {
	Vcode    string `path:"vcode"`
	Buid     int    `form:"buid"`
	ClientID string `form:"client_id"`
}

type VerifyResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type CheckRequest struct {
}

type CheckResponse struct {
	Buid int `json:"buid"`
}

type AddKeyRequest struct {
}

type AddKeyResponse struct {
	Key string `json:"key"`
}

type DeleteKeyRequest struct {
	Key string `form:"key"`
}

type DeleteKeyResponse struct {
}

type GetKeyListRequest struct {
}

type GetKeyListResponse struct {
	Balance int      `json:"balance"`
	Keys    []string `json:"keys"`
}

type RechargeRequest struct {
	Buid   int `json:"buid"`
	Amount int `json:"amount"`
}

type RechargeResponse struct {
	Ok bool `json:"ok"`
}
