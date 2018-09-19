package go_umeng

import (
	"strconv"
	"time"
)

type UnicastIOS struct {
	NotificationIOS
}

func NewUnicastIOS(textMessage string, deviceToken string,account UmengAccount) {
	unicast := &UnicastIOS{
		NotificationIOS{},
	}
	unicast.setConfig(account.APP_MASTER_SECRET)

	payload := &PayloadIOS{}
	aps := &PayloadIOSAps{
		Alert: textMessage,
	}
	payload.Aps = aps

	data := &UmengNotificationData{}
	data.Payload = payload
	data.AppKey = account.APP_KEY
	data.DeviceTokens = deviceToken
	data.ProductionMode = true
	data.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data.Type = Unicast
	unicast.data = data
	unicast.send()
}

//根据别名发送
func SendIOSCustomizedcast(alias string,aliasType string,msg string,account UmengAccount)  {
	unicast := &UnicastIOS{
		NotificationIOS{},
	}
	unicast.setConfig(account.APP_MASTER_SECRET)

	payload := &PayloadIOS{}
	aps := &PayloadIOSAps{
		Alert: msg,
	}
	payload.Aps = aps

	data := &UmengNotificationData{}
	data.Payload = payload
	data.AppKey = account.APP_KEY
	data.Alias=alias
	data.AliasType=aliasType
	data.ProductionMode = true
	data.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data.Type = Customizedcast
	unicast.data = data
	unicast.send()
}
