package wechat

import (
	"fmt"
	"time"
)

/**
 * JsApiTicket数据模型
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type JsApiTicketModel struct {
	Token   string
	Expires *time.Time
	Check   interface{}
	Save    interface{}
}

// Config 返回给用户jssdk配置信息
type JsSdk struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

/**
 * 获取GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetJsApiTicket() string {
	/* 判断是否有自定义方法，用于从mem cache、redis、db中获取token*/
	if model.JsApiTicket.Check != nil {
		model.JsApiTicket.Check.(func(JsApiTicket *JsApiTicketModel))(&model.JsApiTicket)
	}

	/*判断 token 是否存在，或者是否超时*/
	if model.JsApiTicket.Token == "" || model.JsApiTicket.Expires == nil || model.JsApiTicket.Expires.Before(time.Now().UTC()) {
		model.GetJsApiTicketFormWx()
	}

	/*返回token*/
	return model.JsApiTicket.Token
}

/**
 * 调用微信接口获取GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetJsApiTicketFormWx() {

	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiGetJsApiTicket,
		"url_params": map[string]interface{}{
			"access_token": model.GetAccessToken(),
			"type":         "jsapi",
		},
	})

	/*处理请求错误*/
	if err != nil {
		panic(err)
	}

	/*处理返回结果错误*/
	if response.(map[string]interface{})["errcode"] != nil && fmt.Sprint(response.(map[string]interface{})["errcode"]) != "0" {
		panic(response)
	}

	/*处理正常结果*/
	model.JsApiTicket.Token = response.(map[string]interface{})["ticket"].(string)
	expiresIn := response.(map[string]interface{})["expires_in"].(float64)
	expiresTime := time.Now().Add(+time.Second * time.Duration(expiresIn)).UTC()
	model.JsApiTicket.Expires = &expiresTime

	/* 判断是否有自定义方法，用于把Token保存到mem cache、redis、db等*/
	if model.JsApiTicket.Save != nil {
		model.JsApiTicket.Save.(func(model *JsApiTicketModel))(&model.JsApiTicket)
	}
}

/**
 * 获取JsSdk
 * @author ChengCheng
 * @param url 当前网页的URL，不包含#及其后面部分
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetJsSdk(url string) (jsSdk JsSdk) {
	/* 获取要签名的相关参数*/
	ticketStr := model.GetJsApiTicket()
	nonceStr := RandomStr(16)
	timestamp := time.Now().Unix()

	/*签名*/
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, url)
	sigStr := Signature(str)

	/*jsSdk*/
	jsSdk.AppID = model.Config.AppId
	jsSdk.NonceStr = nonceStr
	jsSdk.Timestamp = timestamp
	jsSdk.Signature = sigStr
	return jsSdk
}
