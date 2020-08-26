package wechat

import "encoding/xml"

// CommonToken 消息中通用的结构
type BaseMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

//EventPic 发图事件推送
type EventPic struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

//MsgDevice 设备消息响应
type MsgDevice struct {
	DeviceType string
	DeviceID   string
	SessionID  string
	OpenID     string
}

//MixMessage 存放所有微信发送过来的消息和事件
type MixMessage struct {
	BaseMsg

	//基本消息
	MsgID         int64   `xml:"MsgId"` //其他消息推送过来是MsgId
	TemplateMsgID int64   `xml:"MsgID"` //模板消息推送成功的消息是MsgID
	Content       string  `xml:"Content"`
	Recognition   string  `xml:"Recognition"`
	PicURL        string  `xml:"PicUrl"`
	MediaID       string  `xml:"MediaId"`
	Format        string  `xml:"Format"`
	ThumbMediaID  string  `xml:"ThumbMediaId"`
	LocationX     float64 `xml:"Location_X"`
	LocationY     float64 `xml:"Location_Y"`
	Scale         float64 `xml:"Scale"`
	Label         string  `xml:"Label"`
	Title         string  `xml:"Title"`
	Description   string  `xml:"Description"`
	URL           string  `xml:"Url"`

	//事件相关
	Event       string `xml:"Event"`
	EventKey    string `xml:"EventKey"`
	Ticket      string `xml:"Ticket"`
	Latitude    string `xml:"Latitude"`
	Longitude   string `xml:"Longitude"`
	Precision   string `xml:"Precision"`
	MenuID      string `xml:"MenuId"`
	Status      string `xml:"Status"`
	SessionFrom string `xml:"SessionFrom"`

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`

	SendPicsInfo struct {
		Count   int32      `xml:"Count"`
		PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	}

	// 第三方平台相关
	InfoType                     string `xml:"InfoType"`
	AppID                        string `xml:"AppId"`
	ComponentVerifyTicket        string `xml:"ComponentVerifyTicket"`
	AuthorizerAppd               string `xml:"AuthorizerAppid"`
	AuthorizationCode            string `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime int64  `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string `xml:"PreAuthCode"`

	// 卡券相关
	CardID              string `xml:"CardId"`
	RefuseReason        string `xml:"RefuseReason"`
	IsGiveByFriend      int32  `xml:"IsGiveByFriend"`
	FriendUserName      string `xml:"FriendUserName"`
	UserCardCode        string `xml:"UserCardCode"`
	OldUserCardCode     string `xml:"OldUserCardCode"`
	OuterStr            string `xml:"OuterStr"`
	IsRestoreMemberCard int32  `xml:"IsRestoreMemberCard"`
	UnionID             string `xml:"UnionId"`

	// 内容审核相关
	IsRisky       bool   `xml:"isrisky"`
	ExtraInfoJSON string `xml:"extra_info_json"`
	TraceID       string `xml:"trace_id"`
	StatusCode    int    `xml:"status_code"`

	//设备相关
	MsgDevice
}

//Reply 消息回复
type MixReply struct {
	MsgType string
	MsgData interface{}
}

//ReplyMusic 消息回复
type ReplyMusic struct {
	BaseMsg
	Music struct {
		Title        string `xml:"Title"        `
		Description  string `xml:"Description"  `
		MusicUrL     string `xml:"MusicUrl"     `
		HQMusicUrl   string `xml:"HQMusicUrl"   `
		ThumbMediaId string `xml:"ThumbMediaId"`
	} `xml:"Music"`
}

//ReplyNews  消息回复
type ReplyNews struct {
	BaseMsg
	ArticleCount int        `xml:"ArticleCount"`
	Articles     []*Article `xml:"Articles>item,omitempty"`
}

//Article 单篇文章
type Article struct {
	Title       string `xml:"Title,omitempty"`
	Description string `xml:"Description,omitempty"`
	PicURL      string `xml:"PicUrl,omitempty"`
	URL         string `xml:"Url,omitempty"`
}

//Text 文本消息
type ReplyText struct {
	BaseMsg
	Content string `xml:"Content"`
}

//Video 视频消息
type ReplyVideo struct {
	BaseMsg
	Video struct {
		MediaID     string `xml:"MediaId"`
		Title       string `xml:"Title,omitempty"`
		Description string `xml:"Description,omitempty"`
	} `xml:"Video"`
}

//Voice 语音消息
type ReplyVoice struct {
	BaseMsg
	Voice struct {
		MediaID string `xml:"MediaId"`
	} `xml:"Voice"`
}
