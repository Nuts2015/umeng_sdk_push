package umeng_sdk_push

/*
 *全局配置
 */
const (
	HOST        = "http://msg.umeng.com"
	UPLOAD_PATH = "/upload"
	POST_PATH   = "/api/send"
)

//after open
const (
	GO_APP      = "go_app"
	GO_URL      = "go_url"
	GO_ACTIVITY = "go_activity"
	GO_CUSTOM   = "go_custom"
)

//display payload type
const (
	NOTIFICATION = "notification"
	MESSAGE      = "message"
)

//os type 手机类型 1.苹果手机 2.安卓手机
const (
	IOS     = 1
	ANDROID = 2
)

//message type 消息类型，如定向发送
const (
	Unicast        = "unicast"
	Broadcast      = "broadcast"
	Customizedcast = "customizedcast"
	Filecast       = "filecast"
	Groupcast      = "groupcast"
)

type UmengAccount struct {
	APP_KEY           string
	APP_MASTER_SECRET string
}
