package main

import (
	"github.com/zhuxiujia/umeng-sdk-push"
	"fmt"
)

func main() {
	os := 1
	if os == umeng_sdk_push.IOS {
		resutlt, error := umeng_sdk_push.NewUnicastIOS("恭喜你，你的作品被添加到首页！", "f39a628311ffe034ebe9be661315e76382672d770b9a15667fdb1fbac9668cce", umeng_sdk_push.UmengAccount{
			APP_KEY:           "5a4aeb8d8f4a9d6d5a0000fd",
			APP_MASTER_SECRET: "jib7uvmpbrvxmmfc1dmi6xfqj8xtdmud",
		})
		//resutlt, error := umeng_sdk_push.SendIOSCustomizedcast("18969542172", "phone", "恭喜你，你的作品被添加到首页！", umeng_sdk_push.UmengAccount{
		//	APP_KEY:           "5a4aeb8d8f4a9d6d5a0000fd",
		//	APP_MASTER_SECRET: "jib7uvmpbrvxmmfc1dmi6xfqj8xtdmud",
		//})
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println(resutlt)
	} else if os == umeng_sdk_push.ANDROID {
		//resutlt, error :=umeng_sdk_push.NewAndroidUnicast("借我说得出口的旦旦誓言", "AjfquSK5oypIYnr3HnYYD2aEjxrIQapcJrAmMF95CRZe", umeng_sdk_push.UmengAccount{
		//	APP_KEY:           "5a4aeb8d8f4a9d6d5a0000fd",
		//	APP_MASTER_SECRET: "jib7uvmpbrvxmmfc1dmi6xfqj8xtdmud",
		//})
		resutlt, error := umeng_sdk_push.NewAndroidCustomizedcast("18969542172", "phone", "恭喜你，你的作品被添加到首页！", umeng_sdk_push.UmengAccount{
			APP_KEY:           "5a4aeb8d8f4a9d6d5a0000fd",
			APP_MASTER_SECRET: "jib7uvmpbrvxmmfc1dmi6xfqj8xtdmud",
		})
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println(resutlt)
	}
}
