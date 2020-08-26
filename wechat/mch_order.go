package wechat

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
)

/**
 * 下单接口返回的结果
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type PreOrderResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppID      string `xml:"appid,omitempty"`
	MchID      string `xml:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	TradeType  string `xml:"trade_type,omitempty"`
	PrePayID   string `xml:"prepay_id,omitempty"`
	CodeURL    string `xml:"code_url,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
}

/**
 * Order
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type Order struct {
	AppId          string   `xml:"appid"`                 //微信支付分配的公众账号ID（企业号corpid即为此appId）
	MchId          string   `xml:"mch_id"`                // 微信支付分配的商户号
	DeviceInfo     string   `xml:"device_info,omitempty"` //自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	NonceStr       string   `xml:"nonce_str"`             //随机字符串，长度要求在32位以内
	Sign           string   `xml:"sign"`                  //通过签名算法计算得出的签名值
	SignType       string   `xml:"sign_type,omitempty"`   //签名类型，默认为MD5，支持HMAC-SHA256和MD5
	Body           string   `xml:"body"`                  //商品简单描述，该字段请按照规范传递
	Detail         string   `xml:"detail,omitempty"`      //商品详细描述
	Attach         string   `xml:"attach,omitempty"`      // 附加数据
	OutTradeNo     string   `xml:"out_trade_no"`          // 商户订单号
	FeeType        string   `xml:"fee_type,omitempty"`    // 标价币种
	TotalFee       string   `xml:"total_fee"`             // 标价金额
	SpbillCreateIP string   `xml:"spbill_create_ip"`      // 终端IP
	TimeStart      string   `xml:"time_start,omitempty"`  // 交易起始时间
	TimeExpire     string   `xml:"time_expire,omitempty"` // 交易结束时间
	GoodsTag       string   `xml:"goods_tag,omitempty"`   // 订单优惠标记
	NotifyURL      string   `xml:"notify_url"`            // 通知地址
	TradeType      string   `xml:"trade_type"`            // 交易类型
	ProductID      string   `xml:"product_id,omitempty"`  // 商品ID
	LimitPay       string   `xml:"limit_pay,omitempty"`   // 上传此参数no_credit--可限制用户不能使用信用卡支付
	OpenId         string   `xml:"openid,omitempty"`      // 用户标识
	SceneInfo      string   `xml:"scene_info,omitempty"`  // 场景信息,该字段为JSON对象数据
	XMLName        struct{} `xml:"xml"`
}

/**
 * 填充order的数据
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (order *Order) LoadOrder(model *MpModel) Order {

	/* 填充 AppId 等*/
	if order.AppId == "" {
		order.AppId = model.Config.AppId
	}
	if order.MchId == "" {
		order.MchId = model.Config.MchId
	}
	if order.NonceStr == "" {
		order.NonceStr = RandomStr(32)
	}
	if order.SpbillCreateIP == "" {
		order.SpbillCreateIP = "0.0.0.0"
	}
	if order.TradeType == "" {
		order.TradeType = "JSAPI"
	}
	if order.NotifyURL == "" {
		order.NotifyURL = model.Config.MchNotifyURL
	}
	if order.Body == "" {
		order.Body = "wx pay"
	}

	/*签名*/
	order.Sign, _ = ParamSign(map[string]string{
		"appid":            order.AppId,
		"mch_id":           order.MchId,
		"device_info":      order.DeviceInfo,
		"nonce_str":        order.NonceStr,
		"sign_type":        order.SignType,
		"body":             order.Body,
		"detail":           order.Detail,
		"attach":           order.Attach,
		"out_trade_no":     order.OutTradeNo,
		"fee_type":         order.FeeType,
		"total_fee":        order.TotalFee,
		"spbill_create_ip": order.SpbillCreateIP,
		"time_start":       order.TimeStart,
		"time_expire":      order.TimeExpire,
		"goods_tag":        order.GoodsTag,
		"notify_url":       order.NotifyURL,
		"trade_type":       order.TradeType,
		"product_id":       order.ProductID,
		"limit_pay":        order.LimitPay,
		"openid":           order.OpenId,
		"scene_info":       order.SceneInfo,
	}, model.Config.MchSecret)

	/*返回order信息*/
	return *order
}

/**
 * js调起支付的参数
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type JsPay struct {
	AppId     string `json:"appId"`
	Timestamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	SignType  string `json:"signType"`
	Package   string `json:"package"`
	PaySign   string `json:"paySign"`
}

/**
 * PaidResult 下单回调
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type PaidResult struct {
	ReturnCode *string `xml:"return_code"`
	ReturnMsg  *string `xml:"return_msg"`

	AppID              *string `xml:"appid" json:"appid"`
	MchID              *string `xml:"mch_id"`
	DeviceInfo         *string `xml:"device_info"`
	NonceStr           *string `xml:"nonce_str"`
	Sign               *string `xml:"sign"`
	SignType           *string `xml:"sign_type"`
	ResultCode         *string `xml:"result_code"`
	ErrCode            *string `xml:"err_code"`
	ErrCodeDes         *string `xml:"err_code_des"`
	OpenID             *string `xml:"openid"`
	IsSubscribe        *string `xml:"is_subscribe"`
	TradeType          *string `xml:"trade_type"`
	BankType           *string `xml:"bank_type"`
	TotalFee           *int    `xml:"total_fee"`
	SettlementTotalFee *int    `xml:"settlement_total_fee"`
	FeeType            *string `xml:"fee_type"`
	CashFee            *string `xml:"cash_fee"`
	CashFeeType        *string `xml:"cash_fee_type"`
	CouponFee          *int    `xml:"coupon_fee"`
	CouponCount        *int    `xml:"coupon_count"`

	// coupon_type_$n 这里只声明 3 个，如果有更多的可以自己组合
	CouponType0 *string `xml:"coupon_type_0"`
	CouponType1 *string `xml:"coupon_type_1"`
	CouponType2 *string `xml:"coupon_type_2"`
	CouponID0   *string `xml:"coupon_id_0"`
	CouponID1   *string `xml:"coupon_id_1"`
	CouponID2   *string `xml:"coupon_id_2"`
	CouponFeed0 *string `xml:"coupon_fee_0"`
	CouponFeed1 *string `xml:"coupon_fee_1"`
	CouponFeed2 *string `xml:"coupon_fee_2"`

	TransactionID *string `xml:"transaction_id"`
	OutTradeNo    *string `xml:"out_trade_no"`
	Attach        *string `xml:"attach"`
	TimeEnd       *string `xml:"time_end"`
}

/**
 * PaidResp 消息通知返回
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
type PaidResp struct {
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	XMLName    struct{} `xml:"xml"`
}

/**
 * 调用微信接口获取GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) UnifiedOrder(order Order) (res PreOrderResponse, err error) {

	/*格式化订单数据*/
	order = order.LoadOrder(model)

	/*转换成xml*/
	xmlBody, err := xml.Marshal(order)
	if err != nil {
		return res, err
	}

	/*请求微信接口、获取token*/
	response, err := Request(map[string]interface{}{
		"method": "POST",
		"url":    model.GetApiMchHostUrl() + ApiMchUnifiedOrder,
		"headers": map[string]interface{}{
			"Content-Type": "application/xml",
		},
		"body_params": xmlBody,
		"success": func(response *http.Response) (res interface{}, err error) {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return
			}
			_ = response.Body.Close()
			var orderResponse PreOrderResponse
			err = xml.Unmarshal(body, &orderResponse)
			return orderResponse, err
		},
	})
	if err != nil {
		return res, err
	}

	/* 判断返回结果 */
	preOrderResponse := response.(PreOrderResponse)
	if preOrderResponse.ReturnCode == "SUCCESS" {
		if preOrderResponse.ResultCode == "SUCCESS" {
			return preOrderResponse, nil
		}
		return preOrderResponse, errors.New(preOrderResponse.ErrCode + preOrderResponse.ErrCodeDes)
	}

	/*返回结果*/
	return preOrderResponse, errors.New(preOrderResponse.ReturnCode)
}

/**
 * 调用微信接口获取GetAccessToken
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) GetJsPay(prepayId string) (res JsPay) {
	/*字段*/
	res.AppId = model.Config.AppId
	res.NonceStr = RandomStr(32)
	res.Package = "prepay_id=" + prepayId
	res.Timestamp = fmt.Sprint(time.Now().Unix())
	res.SignType = SignTypeMD5

	/*签名*/
	res.PaySign, _ = ParamSign(map[string]string{
		"appId":     res.AppId,
		"timeStamp": res.Timestamp,
		"nonceStr":  res.NonceStr,
		"package":   res.Package,
		"signType":  res.SignType,
	}, model.Config.MchSecret)
	return res
}

/**
 * PaidVerifySign 支付成功结果验签
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func (model *MpModel) PaidVerifySign(notifyRes PaidResult) bool {
	// STEP1, 转换 struct 为 map，并对 map keys 做排序
	resMap := structs.Map(notifyRes)

	sortedKeys := make([]string, 0, len(resMap))
	for k := range resMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	// STEP2, 对key=value的键值对用&连接起来，略过空值 & sign
	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", cast.ToString(resMap[k]))
		if value != "" && strings.ToLower(k) != "sign" {
			signStrings = signStrings + getTagKeyName(k, &notifyRes) + "=" + value + "&"
		}
	}

	// STEP3, 在键值对的最后加上key=API_KEY
	signStrings = signStrings + "key=" + model.Config.MchSecret

	// STEP4, 根据SignType计算出签名
	var signType string
	if notifyRes.SignType != nil {
		signType = *notifyRes.SignType
	}
	sign, err := CalculateSign(signStrings, signType, model.Config.MchSecret)
	if err != nil {
		return false
	}
	if sign != *notifyRes.Sign {
		return false
	}
	return true
}

/**
 * getTagKeyName
 * @author ChengCheng
 * @date 2020-04-14 22:10:18
 */
func getTagKeyName(key string, notifyRes *PaidResult) string {
	s := reflect.TypeOf(notifyRes).Elem()
	f, _ := s.FieldByName(key)
	name := f.Tag.Get("xml")
	return name
}
