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

	/*处理返回结果错误*/
	if response.(map[string]interface{})["errcode"] != nil {
		panic(response)
	}

	return response
}

/**
 * 拼接AccessToken接口的url
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetQueryMenuUrl() string {
	return model.GetApiHostUrl() + ApiGetMenu
}
