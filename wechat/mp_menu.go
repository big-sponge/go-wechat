package wechat

import (
	"encoding/json"
)

//Button 菜单按钮
type Button struct {
	Type       string    `json:"type,omitempty"`
	Name       string    `json:"name,omitempty"`
	Key        string    `json:"key,omitempty"`
	URL        string    `json:"url,omitempty"`
	MediaID    string    `json:"media_id,omitempty"`
	AppID      string    `json:"appid,omitempty"`
	PagePath   string    `json:"pagepath,omitempty"`
	SubButtons []*Button `json:"sub_button,omitempty"`
}

//SetSubButton 设置二级菜单
func (btn *Button) SetSubButton(name string, subButtons []*Button) {
	btn.Name = name
	btn.SubButtons = subButtons
	btn.Type = ""
	btn.Key = ""
	btn.URL = ""
	btn.MediaID = ""
}

//SetClickButton btn 为click类型
func (btn *Button) SetClickButton(name, key string) {
	btn.Type = "click"
	btn.Name = name
	btn.Key = key
	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetViewButton view类型
func (btn *Button) SetViewButton(name, url string) {
	btn.Type = "view"
	btn.Name = name
	btn.URL = url
	btn.Key = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

// SetScanCodePushButton 扫码推事件
func (btn *Button) SetScanCodePushButton(name, key string) {
	btn.Type = "scancode_push"
	btn.Name = name
	btn.Key = key
	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetScanCodeWaitMsgButton 设置 扫码推事件且弹出"消息接收中"提示框
func (btn *Button) SetScanCodeWaitMsgButton(name, key string) {
	btn.Type = "scancode_waitmsg"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetPicSysPhotoButton 设置弹出系统拍照发图按钮
func (btn *Button) SetPicSysPhotoButton(name, key string) {
	btn.Type = "pic_sysphoto"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetPicPhotoOrAlbumButton 设置弹出拍照或者相册发图类型按钮
func (btn *Button) SetPicPhotoOrAlbumButton(name, key string) {
	btn.Type = "pic_photo_or_album"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

// SetPicWeixinButton 设置弹出微信相册发图器类型按钮
func (btn *Button) SetPicWeixinButton(name, key string) {
	btn.Type = "pic_weixin"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

// SetLocationSelectButton 设置 弹出地理位置选择器 类型按钮
func (btn *Button) SetLocationSelectButton(name, key string) {
	btn.Type = "location_select"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetMediaIDButton  设置 下发消息(除文本消息) 类型按钮
func (btn *Button) SetMediaIDButton(name, mediaID string) {
	btn.Type = "media_id"
	btn.Name = name
	btn.MediaID = mediaID

	btn.Key = ""
	btn.URL = ""
	btn.SubButtons = nil
}

//SetViewLimitedButton  设置 跳转图文消息URL 类型按钮
func (btn *Button) SetViewLimitedButton(name, mediaID string) {
	btn.Type = "view_limited"
	btn.Name = name
	btn.MediaID = mediaID

	btn.Key = ""
	btn.URL = ""
	btn.SubButtons = nil
}

//SetMiniprogramButton  设置 跳转小程序 类型按钮 (公众号后台必须已经关联小程序)
func (btn *Button) SetMiniprogramButton(name, url, appID, pagePath string) {
	btn.Type = "miniprogram"
	btn.Name = name
	btn.URL = url
	btn.AppID = appID
	btn.PagePath = pagePath

	btn.Key = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//reqMenu 设置菜单请求数据
type reqMenu struct {
	Button    []*Button  `json:"button,omitempty"`
	MatchRule *MatchRule `json:"matchrule,omitempty"`
}

//reqDeleteConditional 删除个性化菜单请求数据
type reqDeleteConditional struct {
	MenuID int64 `json:"menuid"`
}

//reqMenuTryMatch 菜单匹配请求
type reqMenuTryMatch struct {
	UserID string `json:"user_id"`
}

//resConditionalMenu 个性化菜单返回结果
type resConditionalMenu struct {
	Button    []Button  `json:"button"`
	MatchRule MatchRule `json:"matchrule"`
	MenuID    int64     `json:"menuid"`
}

//resMenuTryMatch 菜单匹配请求结果
type resMenuTryMatch struct {
	CommonError
	Button []Button `json:"button"`
}

//ResMenu 查询菜单的返回数据
type ResMenu struct {
	CommonError
	Menu struct {
		Button []Button `json:"button"`
		MenuID int64    `json:"menuid"`
	} `json:"menu"`
	Conditionalmenu []resConditionalMenu `json:"conditionalmenu"`
}

//ResSelfMenuInfo 自定义菜单配置返回结果
type ResSelfMenuInfo struct {
	CommonError
	IsMenuOpen   int32 `json:"is_menu_open"`
	SelfMenuInfo struct {
		Button []SelfMenuButton `json:"button"`
	} `json:"selfmenu_info"`
}

//SelfMenuButton 自定义菜单配置详情
type SelfMenuButton struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	URL       string `json:"url,omitempty"`
	Value     string `json:"value,omitempty"`
	SubButton struct {
		List []SelfMenuButton `json:"list"`
	} `json:"sub_button,omitempty"`
	NewsInfo struct {
		List []ButtonNew `json:"list"`
	} `json:"news_info,omitempty"`
}

//ButtonNew 图文消息菜单
type ButtonNew struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Digest     string `json:"digest"`
	ShowCover  int32  `json:"show_cover"`
	CoverURL   string `json:"cover_url"`
	ContentURL string `json:"content_url"`
	SourceURL  string `json:"source_url"`
}

//MatchRule 个性化菜单规则
type MatchRule struct {
	GroupID            string `json:"group_id,omitempty"`
	Sex                string `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
}

/**
 * menu/get
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetMenu() interface{} {
	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiMenuGet,
		"url_params": map[string]interface{}{
			"access_token": model.GetAccessToken(),
		},
	})
	/*处理请求错误*/
	if err != nil {
		panic(err)
	}
	return response
}

/**
 * menu/create
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) CreateMenu(menuJson string) interface{} {
	/*请求微信接口、获取token*/
	var param interface{}
	_ = json.Unmarshal([]byte(menuJson), &param)
	response, err := Request(map[string]interface{}{
		"method":      "POST",
		"url":         model.GetApiHostUrl() + ApiMenuCreate + "?access_token=" + model.GetAccessToken(),
		"body_params": param,
	})
	/*处理请求错误*/
	if err != nil {
		panic(err)
	}
	return response
}

/**
 * menu/delete
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) DeleteMenu() interface{} {
	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiMenuDelete + "?access_token=" + model.GetAccessToken(),
	})
	/*处理请求错误*/
	if err != nil {
		panic(err)
	}
	return response
}

/**
 * menu/addconditional
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) AddConditionalMenu(menuJson string) interface{} {
	/*请求微信接口、获取token*/
	var param interface{}
	_ = json.Unmarshal([]byte(menuJson), &param)
	response, err := Request(map[string]interface{}{
		"method":      "POST",
		"url":         model.GetApiHostUrl() + ApiMenuAddConditional + "?access_token=" + model.GetAccessToken(),
		"body_params": param,
	})
	/*处理请求错误*/
	if err != nil {
		panic(err)
	}
	return response
}

/**
 * menu/delconditional
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) DelConditionalMenu(menuId string) interface{} {
	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"method": "POST",
		"url":    model.GetApiHostUrl() + ApiMenuDelConditional + "?access_token=" + model.GetAccessToken(),
		"body_params": map[string]interface{}{
			"menuid": menuId,
		},
	})
	/*处理请求错误*/
	if err != nil {
		panic(err)
	}
	return response
}

/**
 * menu/trymatch
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) TryMatchMenu(userId string) interface{} {
	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiMenuTryMatch + "?access_token=" + model.GetAccessToken(),
		"body_params": map[string]interface{}{
			"user_id": userId,
		},
	})
	/*处理请求错误*/
	if err != nil {
		panic(err)
	}
	return response
}

/**
 * get_current_selfmenu_info
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetCurrentSelfMenuInfo() interface{} {
	/*请求微信接口*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiMenuGetCurrentSelfMenuInfo + "?access_token=" + model.GetAccessToken(),
	})
	/*处理请求错误*/
	if err != nil {
		panic(err)
	}
	return response
}
