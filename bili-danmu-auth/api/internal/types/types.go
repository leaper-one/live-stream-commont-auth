// Code generated by goctl. DO NOT EDIT.
package types

type ApplyNewVCodeRequest struct {
	Buid     int    `path:"buid"`
	Key      string `form:"key"` // 0 为 devloper 登录, 不可以使用
	ClientID string `form:"client_id"`
}

type ApplyNewVCodeResponse struct {
	Vcode string `json:"vcode"`
}

type AddOneRequest struct {
	Vcode string `path:"vcode"`
	Buid  int    `form:"buid"`
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
	Buid string `json:"buid"`
}

type AddKeyRequest struct {
	Key string `form:"key"`
}

type AddKeyResponse struct {
	Ok bool `json:"ok"`
}

type DeleteKeyRequest struct {
	Key string `form:"key"`
}

type DeleteKeyResponse struct {
	Ok bool `json:"ok"`
}

type GetKeyListRequest struct {
}

type GetKeyListResponse struct {
	Balance int64    `json:"balance"`
	Keys    []string `json:"keys"`
}

type RechargeRequest struct {
	Buid   string `json:"buid"`
	Amount int64  `json:"amount"`
}

type RechargeResponse struct {
	Ok bool `json:"ok"`
}
