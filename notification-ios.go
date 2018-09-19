package umeng_sdk_push

type PayloadIOS struct {
	Aps *PayloadIOSAps `json:"aps"`
}

type PayloadIOSAps struct {
	Alert string `json:"alert"`
}

type NotificationIOS struct {
	UmengNotification
}

func (this *NotificationIOS) setConfig(secret string) {
	this.appMasterSecret = secret
	this.uploadPath = UPLOAD_PATH
	this.postPath = POST_PATH
	this.host = HOST
}
