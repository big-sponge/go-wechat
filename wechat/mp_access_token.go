package wechat

import (
	"time"
)

/**
 * AccessToken数据模型
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type MpAccessTokenModel struct {
	Token   string
	Expires *time.Time
	Check   interface{}
	Save    interface{}
}

/**
 * 获取GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetAccessToken() string {
	/* 判断是否有自定义方法，用于从mem cache、redis、db中获取token*/
	if model.AccessToken.Check != nil {
		model.AccessToken.Check.(func(accessToken *MpAccessTokenModel))(&model.AccessToken)
	}

	/*判断 token 是否存在，或者是否超时*/
	if model.AccessToken.Token == "" || model.AccessToken.Expires == nil || model.AccessToken.Expires.Before(time.Now().UTC()) {
		model.GetAccessTokenFormWx()
	}

	/*返回token*/
	return model.AccessToken.Token
}

/**
 * 调用微信接口获取GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetAccessTokenFormWx() {

	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetAccessTokenUrl(),
		"url_params": map[string]interface{}{
			"grant_type": "client_credential",
			"appid":      model.Config.AppId,
			"secret":     model.Config.AppSecret,
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

	/*处理正常结果*/
	model.AccessToken.Token = response.(map[string]interface{})["access_token"].(string)
	expiresIn := response.(map[string]interface{})["expires_in"].(float64)
	expiresTime := time.Now().Add(+time.Second * time.Duration(expiresIn)).UTC()
	model.AccessToken.Expires = &expiresTime

	/* 判断是否有自定义方法，用于把Token保存到mem cache、redis、db等*/
	if model.AccessToken.Save != nil {
		model.AccessToken.Save.(func(model *MpAccessTokenModel))(&model.AccessToken)
	}
}

/**
 * 拼接AccessToken接口的url
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetAccessTokenUrl() string {
	return model.GetApiHostUrl() + ApiGetAccessToken
}
