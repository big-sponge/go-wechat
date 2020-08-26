package wechat

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
 * GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetAccessToken() (res string) {
	return mp.GetAccessToken()
}

/**
 * ApiDomainIp
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetApiDomainIp() (res interface{}, err error) {
	return mp.GetApiDomainIp()
}

/**
 * 网络检查
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func Check(action string, checkOperator string) (res interface{}, err error) {
	return mp.Check(action, checkOperator)
}

/**
 * 获取菜单
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetOauthUserInfo(code string) (mpOauthUserInfo MpOauthUserInfo, err error) {
	return mp.GetOauthUserInfo(code)
}


/**
 * 统一下单
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func UnifiedOrder(order Order) (res PreOrderResponse, err error) {
	return mp.UnifiedOrder(order)
}

/**
 * 验证版本
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func PaidVerifySign(notifyRes PaidResult) bool {
	return mp.PaidVerifySign(notifyRes)
}

/**
 * 查询订单
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func OrderQuery(order Order) (res PreOrderResponse, err error) {
	return
}

/**
 * 获取js支付
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetJsPay(prepayId string) (res JsPay) {
	return mp.GetJsPay(prepayId)
}

/**
 * GetJsApiTicket
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetJsApiTicket() (res string) {
	return mp.GetJsApiTicket()
}

/**
 * GetJsSdk
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetJsSdk(url string) (res JsSdk) {
	return mp.GetJsSdk(url)
}

/**
 * 获取菜单
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetMenu() interface{} {
	return mp.GetMenu()
}

/**
 * 删除菜单
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func DeleteMenu() interface{} {
	return mp.DeleteMenu()
}

/**
 * 设置菜单
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func AddConditionalMenu(menuJson string) interface{} {
	return mp.AddConditionalMenu(menuJson)
}

/**
 * 获取菜单
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func GetCurrentSelfMenuInfo() interface{} {
	return mp.GetCurrentSelfMenuInfo()
}

