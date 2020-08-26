package wechat

import (
	"github.com/pkg/errors"
)

/**
 * AccessToken数据模型
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type MpOauthAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

/**
 * userinfo数据模型
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type MpOauthUserInfo struct {
	Openid     string        `json:"openid"`
	Nickname   string        `json:"nickname"`
	Sex        float64       `json:"sex"`
	Province   string        `json:"province"`
	City       string        `json:"city"`
	Country    string        `json:"country"`
	HeadImgurl string        `json:"headimgurl"`
	Privilege  []interface{} `json:"privilege"`
	UnionId    string        `json:"unionid"`
}

/**
 * 调用微信接口获取GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetOauthAccessToken(code string) (oauthAccessToken MpOauthAccessToken, err error) {
	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiGetAuthAccessToken,
		"url_params": map[string]interface{}{
			"appid":      model.Config.AppId,
			"secret":     model.Config.AppSecret,
			"grant_type": "authorization_code",
			"code":       code,
		},
	})
	if err != nil {
		return oauthAccessToken, err
	}

	/* 返回 AccessToken */
	if response.(map[string]interface{})["access_token"] != nil {
		oauthAccessToken.AccessToken = response.(map[string]interface{})["access_token"].(string)
		oauthAccessToken.Openid = response.(map[string]interface{})["openid"].(string)
		return oauthAccessToken, nil
	}

	/* 返回 errmsg */
	if response.(map[string]interface{})["errmsg"] != nil {
		return oauthAccessToken, errors.New(response.(map[string]interface{})["errmsg"].(string))
	}
	return oauthAccessToken, errors.New("GetOauthAccessToken Error.")
}

/**
 * 刷新token
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) OauthRefreshToke(refreshToken string) (oauthAccessToken MpOauthAccessToken, err error) {
	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiAuthRefreshToken,
		"url_params": map[string]interface{}{
			"appid":        model.Config.AppId,
			"grant_type":   "refresh_token",
			"refreshToken": refreshToken,
		},
	})
	if err != nil {
		return oauthAccessToken, err
	}

	/* 返回 AccessToken */
	if response.(map[string]interface{})["access_token"] != nil {
		oauthAccessToken.AccessToken = response.(map[string]interface{})["access_token"].(string)
		oauthAccessToken.Openid = response.(map[string]interface{})["openid"].(string)
		return oauthAccessToken, nil
	}

	/* 返回 errmsg */
	if response.(map[string]interface{})["errmsg"] != nil {
		return oauthAccessToken, errors.New(response.(map[string]interface{})["errmsg"].(string))
	}
	return oauthAccessToken, errors.New("GetOauthAccessToken Error.")
}

/**
 * 检查token
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) ApiAuthCheck(accessToken string) (res map[string]interface{}, err error) {
	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiAuthCheck,
		"url_params": map[string]interface{}{
			"appid":        model.Config.AppId,
			"access_token": accessToken,
		},
	})
	if err != nil {
		return res, err
	}
	return response.(map[string]interface{}), nil
}

/**
 * 拉取用户信息
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetOauthUserInfo(code string) (mpOauthUserInfo MpOauthUserInfo, err error) {

	/* 调用微信接口获取GetAccessToken */
	oauthAccessToken, err := model.GetOauthAccessToken(code)
	if err != nil {
		return mpOauthUserInfo, err
	}

	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiGetAuthUserInfo,
		"url_params": map[string]interface{}{
			"access_token": oauthAccessToken.AccessToken,
			"openid":       oauthAccessToken.Openid,
			"lang":         "zh_CN",
		},
	})
	if err != nil {
		return mpOauthUserInfo, err
	}

	/* 返回 AccessToken */
	if response.(map[string]interface{})["openid"] != nil {
		mpOauthUserInfo.Openid = response.(map[string]interface{})["openid"].(string)
		mpOauthUserInfo.Nickname = response.(map[string]interface{})["nickname"].(string)
		mpOauthUserInfo.Province = response.(map[string]interface{})["province"].(string)
		mpOauthUserInfo.City = response.(map[string]interface{})["city"].(string)
		mpOauthUserInfo.Country = response.(map[string]interface{})["country"].(string)
		mpOauthUserInfo.HeadImgurl = response.(map[string]interface{})["headimgurl"].(string)
		mpOauthUserInfo.Sex = response.(map[string]interface{})["sex"].(float64)
		mpOauthUserInfo.Privilege = response.(map[string]interface{})["privilege"].([]interface{})
		return mpOauthUserInfo, nil
	}

	/* 返回 errmsg */
	if response.(map[string]interface{})["errmsg"] != nil {
		return mpOauthUserInfo, errors.New(response.(map[string]interface{})["errmsg"].(string))
	}
	return mpOauthUserInfo, errors.New("GetOauthUserInfo Error.")
}
