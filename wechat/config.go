package wechat

const (
	ApiHost           = "https://api.weixin.qq.com"
	ApiGetAccessToken = "/cgi-bin/token"

	/*menu*/
	/*查询自定义菜单*/
	ApiGetMenu = "/cgi-bin/get_current_selfmenu_info"
	/*创建自定义菜单*/
	ApiCreateMenu = "/cgi-bin/menu/create"
	/*删除自定义菜单*/
	ApiDeleteMenu = "/cgi-bin/menu/delete"

	/*个性化菜单*/
	/*创建个性化菜单*/
	ApiAddConditional = "/cgi-bin/menu/addconditional"
	/*删除个性化菜单*/
	ApiDelConditional = "/cgi-bin/menu/delconditional"

	/*获取自定义菜单配置*/
	ApiGetAllMenuInfo = "/cgi-bin/menu/get"
)
