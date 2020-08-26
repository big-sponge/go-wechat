package wechat

const (
	/* host */
	ApiHost    = "https://api.weixin.qq.com"     /*微信接口host*/
	ApiMchHost = "https://api.mch.weixin.qq.com" /*微信支付接口host*/

	/* 基础 */
	ApiGetAccessToken      = "/cgi-bin/token"             /*获取token*/
	ApiGetApiDomainIp      = "/cgi-bin/get_api_domain_ip" /*获取ip*/
	ApiGetApiCallbackCheck = "/cgi-bin/callback/check"    /*检查网络*/

	/*菜单*/
	ApiMenuCreate                 = "/cgi-bin/menu/create"               /*创建菜单*/
	ApiMenuGet                    = "/cgi-bin/menu/get"                  /*创建菜单*/
	ApiMenuDelete                 = "/cgi-bin/menu/delete"               /*删除菜单*/
	ApiMenuAddConditional         = "/cgi-bin/menu/addconditional"       /*创建个性化菜单*/
	ApiMenuDelConditional         = "/cgi-bin/menu/delconditional"       /*删除个性化菜单*/
	ApiMenuTryMatch               = "/cgi-bin/menu/trymatch"             /*删除个性化菜单*/
	ApiMenuGetCurrentSelfMenuInfo = "/cgi-bin/get_current_selfmenu_info" /*查询自定义菜单*/

	/*用户*/
	ApiGetAuthAccessToken = "/sns/oauth2/access_token"  /*通过code换取网页授权access_token*/
	ApiAuthRefreshToken   = "/sns/oauth2/refresh_token" /*刷新access_token*/
	ApiAuthCheck          = "/sns/auth"                 /*检验授权凭证（access_token）是否有效*/
	ApiGetAuthUserInfo    = "/sns/userinfo"             /*拉取用户信息*/
	ApiGetJsApiTicket     = "/cgi-bin/ticket/getticket" /*获取ticket*/

	/*模板消息*/
	ApiSendTemplateMessage = "/cgi-bin/message/template/send" /*发送模板消息*/

	/*微信支付*/
	ApiMchUnifiedOrder = "/pay/unifiedorder" /*微信支付-统一下单接口*/
	ApiMchQueryOrder   = "/pay/orderquery"   /*微信支付-查询订单*/
)
