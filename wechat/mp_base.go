package wechat

var (
	mp MpModel
)

type MpModel struct {
	Config      MpConfigModel
	AccessToken MpAccessTokenModel
}

type MpConfigModel struct {
	AppId     string
	AppSecret string
}
