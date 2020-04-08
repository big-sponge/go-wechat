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
	ApiHost     string
}

/**
 * 配置数据模型
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type MpConfigModel struct {
	AppId     string
	AppSecret string
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
