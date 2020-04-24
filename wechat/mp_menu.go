package wechat

import (
	"time"
)

/**
 * AccessToken数据模型
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type MpAccessTokenModel3 struct {
	Token   string
	Expires *time.Time
	Check   interface{}
	Save    interface{}
}

/**
 * 调用微信接口获取GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetMenu() interface{} {

	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetQueryMenuUrl(),
		"url_params": map[string]interface{}{
			"access_token": GetAccessToken(),
		},
	})

	/*处理请求错误*/
	if err != nil {
		panic(err)
	}

	return response
}

/**
 * 拼接获取Menu接口的url
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetQueryMenuUrl() string {
	return model.GetApiHostUrl() + ApiGetMenu
}

//自定义菜单URL
func (model *MpModel) GetMenuUrl(ApiPath string) string {
	return model.GetApiHostUrl() + ApiPath
}

//自定义菜单可以实现多种类型的按钮
const (
	ButtonTypeView        = "view"        // 跳转URL
	ButtonTypeClick       = "click"       // 点击推事件
	ButtonTypeMiniProgram = "miniprogram" // 小程序

	//仅支持微信iPhone5.4.1以上版本，和Android5.4以上版本的微信用户，旧版本微信用户点击后将没有回应，开发者也不能正常接收到事件推送
	ButtonTypeScanCodePush    = "scancode_push"      // 扫码推事件
	ButtonTypeScanCodeWaitMsg = "scancode_waitmsg"   // 扫码带提示
	ButtonTypePicSysPhoto     = "pic_sysphoto"       // 系统拍照发图
	ButtonTypePicPhotoOrAlbum = "pic_photo_or_album" // 拍照或者相册发图
	ButtonTypePicWeixin       = "pic_weixin"         // 微信相册发图
	ButtonTypeLocationSelect  = "location_select"    // 发送位置

	//专门给第三方平台旗下未微信认证（具体而言，是资质认证未通过）的订阅号准备的事件类型，它们是没有事件推送的，能力相对受限，其他类型的公众号不必使用。
	ButtonTypeMediaId     = "media_id"     // 下发消息
	ButtonTypeViewLimited = "view_limited" // 跳转图文消息URL
)

//创建自定义菜单
func (model *MpModel) CreateMenu() interface{} { /*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetQueryMenuUrl(),
		"url_params": map[string]interface{}{
			"access_token": GetAccessToken(),
		},
	})

	/*处理请求错误*/
	if err != nil {
		panic(err)
	}

	response2 := response
	return response2
}

//删除最定义菜单
