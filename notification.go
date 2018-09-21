package umeng_sdk_push

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"log"
)

type UmengNotificationData struct {
	AppKey         string `json:"appkey"`
	Timestamp      string `json:"timestamp"`
	Type           string `json:"type"`
	DeviceTokens   string `json:"device_tokens"`
	Alias          string `json:"alias"`
	AliasType      string `json:"alias_type"`
	FileId         string `json:"file_id"`
	Filter         string `json:"filter"`
	Feedback       string `json:"feedback"`
	Description    string `json:"description"`
	ThirdpartyId   string  `json:"thirdparty_id"`
	ProductionMode bool `json:"production_mode"`
	Payload        interface{} `json:"payload"`
}

type UmengNotification struct {
	host            string
	uploadPath      string
	postPath        string
	appMasterSecret string
	data            *UmengNotificationData
}

func (this *UmengNotification) send() (UmengResult,error) {
	var result UmengResult
	url := this.host + this.postPath
	postBody, err := json.Marshal(this.data)
	if err != nil {
		log.Println(err.Error())
		return result,err
	}
	sign := MD5("POST" + url + string(postBody) + this.appMasterSecret)
	url = url + "?sign=" + sign
	bufReader := bytes.NewReader(postBody)
	resp, err := http.Post(url, "application/json", bufReader)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		log.Println(err.Error())
		return result,err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return result,err
	}
	json.Unmarshal(content,&result)
	return result,nil
}
