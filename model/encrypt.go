package model

type EncryptReq struct {
	DeviceCode string `json:"deviceCode"`
	Content    string `json:"content"`
}

type EncryptResp struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
