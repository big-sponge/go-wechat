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


	wechat.Config(wechat.MpConfigModel{
		AppId:     "wx030ea3d763b7c9fb",
		AppSecret: "91b7a02555480e57703874b1244ccdfc",
	})


	fmt.Println(wechat.Menu())
}
