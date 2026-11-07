// 小程序登录态校验示例：检查 code2Session 得到的 openid / session_key 是否仍然有效。
//
//	export WECHAT_APPID=wx...
//	export WECHAT_APPSECRET=...
//	go run ./miniProgram/checkSessionKey
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

	openID := prompt(reader, "请输入 openid：")
	sessionKey := prompt(reader, "请输入 session_key：")

	if openID == "" || sessionKey == "" {
		fmt.Println("openid 和 session_key 不能为空")
		return
	}

	app, err := miniProgram.New(&miniProgram.Config{
		AppID:     os.Getenv("WECHAT_APPID"),
		AppSecret: os.Getenv("WECHAT_APPSECRET"),
	})
	if err != nil {
		log.Fatalf("初始化失败: %v", err)
	}

	// signature = HMAC-SHA256(session_key, 空消息) 的十六进制，由 SDK 计算。
	signature := miniProgram.CheckSessionSignature(sessionKey)

	res, err := app.CheckSessionKey(context.Background(), openID, signature)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}
	if !res.OK() {
		log.Fatalf(
			"登录态已失效: %v（errcode=%d errmsg=%s）",
			miniProgram.CheckSessionKeyError(res),
			res.ErrCode,
			res.ErrMsg,
		)
	}

	log.Println("登录态有效")
}

func prompt(reader *bufio.Reader, tip string) string {
	fmt.Print(tip)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取失败: %v", err)
	}

	return strings.TrimSpace(line)
}
