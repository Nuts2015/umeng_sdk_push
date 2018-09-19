package umeng_sdk_push

import (
	"strconv"
	"time"
)

type UnicastAndroid struct {
	NotificationAndroid
}

func NewAndroidUnicast(textMessage string, deviceToken string,account UmengAccount) *UnicastAndroid {
	unicast := &UnicastAndroid{
		NotificationAndroid{},
	}
	unicast.setConfig(account.APP_MASTER_SECRET)
	payload := &PayloadAndroid{}
	payload.DisplayType = NOTIFICATION
	payload.Extra = make(map[string]interface{})
	payload.Body = &PayloadBodyAndroid{}
	payload.Body.Title = "消息通知"
	payload.Body.Ticker = "消息通知"
	payload.Body.PlayLights = true
	payload.Body.PlaySound = true
	payload.Body.PlayVibrate = true
	payload.Body.Text = textMessage
	payload.Body.AfterOpen = GO_APP

	data := &UmengNotificationData{}
	data.Payload = payload
	data.AppKey = account.APP_KEY
	data.DeviceTokens = deviceToken
	data.ProductionMode = true
	data.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data.Type = Unicast
	unicast.data = data

	unicast.send()
	return unicast
}

func SendAndroidCustomizedcast(alias string,aliasType string,msg string,account UmengAccount)  *UnicastAndroid{
	unicast := &UnicastAndroid{
		NotificationAndroid{},
	}
	unicast.setConfig(account.APP_MASTER_SECRET)
	payload := &PayloadAndroid{}
	payload.DisplayType = NOTIFICATION
	payload.Extra = make(map[string]interface{})
	payload.Body = &PayloadBodyAndroid{}
	payload.Body.Title = "消息通知"
	payload.Body.Ticker = "消息通知"
	payload.Body.PlayLights = true
	payload.Body.PlaySound = true
	payload.Body.PlayVibrate = true
	payload.Body.Text = msg
	payload.Body.AfterOpen = GO_APP

	data := &UmengNotificationData{}
	data.Payload = payload
	data.AppKey = account.APP_KEY
	data.Alias=alias
	data.AliasType=aliasType
	data.ProductionMode = true
	data.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data.Type = Unicast
	unicast.data = data

	unicast.send()
	return unicast
}