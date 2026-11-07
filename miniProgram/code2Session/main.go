// 小程序登录示例：用 wx.login 得到的 code 换取 openid 与 session_key。
//
//	export WECHAT_APPID=wx...
//	export WECHAT_APPSECRET=...
//	go run ./miniProgram/code2Session -code=0d3d1NFa1YlqKM0koXFa1yGSd22d1NFt
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antsurge/atto-wechat/platform/miniProgram"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入 wx.login 返回的临时 code：")

	code, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败：", err)
		return
	}

	code = strings.TrimSpace(code)

	if code == "" {
		fmt.Println("code 不能为空")
		return
	}

	app, err := miniProgram.New(&miniProgram.Config{
		AppID:     os.Getenv("WECHAT_APPID"),
		AppSecret: os.Getenv("WECHAT_APPSECRET"),
	})
	if err != nil {
		log.Fatalf("初始化失败: %v", err)
	}

	res, err := app.Code2Session(context.Background(), code)
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
