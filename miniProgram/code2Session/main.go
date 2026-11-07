// 小程序登录示例：用 wx.login 得到的 code 换取 openid 与 session_key。
//
//	export WECHAT_APPID=wx...
//	export WECHAT_APPSECRET=...
//	go run ./miniProgram/code2Session -code=0d3d1NFa1YlqKM0koXFa1yGSd22d1NFt
package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/antsurge/atto-wechat/platform/miniProgram"
)

func main() {
	code := flag.String("code", "", "wx.login 返回的临时 code")
	flag.Parse()

	if *code == "" {
		log.Fatal("请通过 -code 传入 wx.login 返回的 code")
	}

	app, err := miniProgram.New(&miniProgram.Config{
		AppID:     os.Getenv("WECHAT_APPID"),
		AppSecret: os.Getenv("WECHAT_APPSECRET"),
	})
	if err != nil {
		log.Fatalf("初始化失败: %v", err)
	}

	res, err := app.Code2Session(context.Background(), *code)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}
	if !res.OK() {
		log.Fatalf("微信拒绝: %v", miniProgram.Code2SessionError(res))
	}

	log.Printf("openid: %s", res.OpenID)
	log.Printf("unionid: %s", res.UnionID)
	log.Printf("session_key: %s", res.SessionKey)
}
