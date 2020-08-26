package wechat


var (
	mp MpModel
)

/**
 * 服务号数据模型
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type MpModel struct {
	Config      MpConfigModel
	AccessToken MpAccessTokenModel
	JsApiTicket JsApiTicketModel
	ApiHost     string
	ApiMchHost  string
}

/**
 * 配置数据模型
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type MpConfigModel struct {
	AppId        string
	AppSecret    string
	MchId        string
	MchSecret    string
	MchNotifyURL string
}

type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

/**
 * 获取HOST
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetApiHostUrl() string {
	if model.ApiHost != "" {
		return model.ApiHost
	}
	return ApiHost
}

/**
 * 获取HOST
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetApiMchHostUrl() string {
	if model.ApiMchHost != "" {
		return model.ApiMchHost
	}
	return ApiMchHost
}

/**
 * 获取微信API接口IP地址
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetApiDomainIp() (res interface{}, err error) {
	response, err := Request(map[string]interface{}{
		"url": model.GetApiHostUrl() + ApiGetApiDomainIp + "?access_token=" + model.GetAccessToken(),
	})
	return response, err
}

/**
 * 获取微信API接口IP地址
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) Check(action string, checkOperator string) (res interface{}, err error) {
	if action == "" {
		action = "all"
	}
	if checkOperator == "" {
		checkOperator = "DEFAULT"
	}
	response, err := Request(map[string]interface{}{
		"method": "POST",
		"url":    model.GetApiHostUrl() + ApiGetApiCallbackCheck + "?access_token=" + model.GetAccessToken(),
		"body_params": map[string]interface{}{
			"action":         action,
			"check_operator": checkOperator,
		},
	})
	return response, err
}
