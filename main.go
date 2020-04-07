package main

import (
	"fmt"
	"github.com/big-sponge/go-wechat/wechat"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Print("主程序异常: ", err, ",Recovery:", wechat.Recovery())
		}
	}()

	fmt.Println(wechat.GetAccessToken())
}
