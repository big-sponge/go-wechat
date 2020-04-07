package wechat

import "fmt"

func GetAccessToken() (res interface{}, err error) {

	mp.AccessToken.Check = func(m *MpAccessTokenModel) {
		m.AccessToken = "1"
	}

	fmt.Println(mp.AccessToken.Get())
	fmt.Println(mp.AccessToken.Get())

	return res, err
}
