package wechat

/**
 * GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetAccessToken() (res string) {
	return mp.GetAccessToken()
}

/**
 * 配置
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func Config(config MpConfigModel) () {
	mp.Config = config
	return
}

/**
 * 获取菜单
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func Menu() interface{} {
	return mp.GetMenu()
}
