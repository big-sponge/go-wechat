package wechat

import (
	"time"
)

var (
	mp MpModel
)

type MpModel struct {
	AccessToken MpAccessTokenModel
}

type MpAccessTokenModel struct {
	AccessToken string
	ExpiresIn   float64
	ExpiresTime *time.Time
	Check       interface{}
	Save        interface{}
}

func (model *MpAccessTokenModel) Get() string {
	/* 判断是否有自定义方法，用于从mem cache、redis、db中获取token*/
	if model.Check != nil {
		model.Check.(func(model *MpAccessTokenModel))(model)
	}

	/*判断 token 是否存在，或者是否超时*/
	if model.AccessToken == "" || model.ExpiresTime == nil || model.ExpiresTime.Before(time.Now().UTC()) {
		model.ReGet()
	}

	/*返回token*/
	return model.AccessToken
}

func (model *MpAccessTokenModel) ReGet() MpAccessTokenModel {

	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx030ea3d763b7c9fb&secret=91b7a02555480e57703874b1244ccdfc",
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
	model.AccessToken = response.(map[string]interface{})["access_token"].(string)
	model.ExpiresIn = response.(map[string]interface{})["expires_in"].(float64)
	expiresTime := time.Now().Add(+time.Second * time.Duration(model.ExpiresIn)).UTC()
	model.ExpiresTime = &expiresTime

	/* 判断是否有自定义方法，用于把Token保存到mem cache、redis、db等*/
	if model.Save != nil {
		model.Save.(func(model *MpAccessTokenModel))(model)
	}
	return *model
}
