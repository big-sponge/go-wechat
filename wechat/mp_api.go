package wechat

func GetAccessToken() (res string) {
	return mp.GetAccessToken()
}

func Config(config MpConfigModel) () {
	mp.Config = config
	return
}

func Menu() interface{} {
	return mp.GetMenu()
}
