package umeng_sdk_push

import (
	"strconv"
	"time"
)

type UnicastIOS struct {
	NotificationIOS
}

func NewUnicastIOS(textMessage string, deviceToken string, account UmengAccount) (UmengResult, error) {
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
	return unicast.send()
}

//根据别名发送
func NewIOSCustomizedcast(alias string, aliasType string, msg string, account UmengAccount) (UmengResult, error) {
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
	data.Alias = alias
	data.AliasType = aliasType
	data.ProductionMode = true
	data.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data.Type = Customizedcast
	unicast.data = data
	return unicast.send()
}

//全部发送
func NewIOSBroadcast(msg string, account UmengAccount) (UmengResult, error) {
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
	data.ProductionMode = true
	data.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data.Type = Broadcast
	unicast.data = data
	return unicast.send()
}
