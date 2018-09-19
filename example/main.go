package main

import "github.com/zhuxiujia/umeng-sdk-push"

func main() {
	os := 1
	if os == umeng_sdk_push.IOS {
		//umeng_sdk_push.NewUnicastIOS("恭喜你，你的作品被添加到首页！", "", umeng_sdk_push.UmengAccount{
		//	APP_KEY:           "",
		//	APP_MASTER_SECRET: "",
		//})
		umeng_sdk_push.SendIOSCustomizedcast("18969542172","phone","恭喜你，你的作品被添加到首页！", umeng_sdk_push.UmengAccount{
			APP_KEY:           "",
			APP_MASTER_SECRET: "",
		})
	} else if os == umeng_sdk_push.ANDROID {
		umeng_sdk_push.NewAndroidUnicast("借我说得出口的旦旦誓言", "AjfquSK5oypIYnr3HnYYD2aEjxrIQapcJrAmMF95CRZe", umeng_sdk_push.UmengAccount{
			APP_KEY:           "",
			APP_MASTER_SECRET: "",
		})
	}
}
