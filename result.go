package umeng_sdk_push

type RetType string

const (
	FAIL    RetType = `FAIL`
	SUCCESS RetType = `SUCCESS`
)

type UmengResult struct {
	Ret  string `json:"ret"`
	Data Data   `json:"data"`
}

type Data struct {
	Error_code string `json:"error_code"`
	Error_msg  string `json:"error_msg"`
	Msg_id     string `json:"msg_id"`
}
