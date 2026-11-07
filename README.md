# atto-wechat-examples

[atto-wechat](https://github.com/antsurge/atto-wechat) 的可运行示例。

## 环境准备

- Go 1.25+
- 微信小程序 AppID / AppSecret（小程序后台「开发管理 → 开发设置」获取）

```bash
export WECHAT_APPID=wx1234567890abcdef
export WECHAT_APPSECRET=0123456789abcdef0123456789abcdef
```

## 示例

### 小程序 code2Session

```bash
go run ./miniProgram/code2Session -code=0d3d1NFa1YlqKM0koXFa1yGSd22d1NFt
```

打印 `openid`、`unionid`、`session_key`。详见 [miniProgram/code2Session](./miniProgram/code2Session)。

其余示例的规划进度见 [PROGRESS.md](./PROGRESS.md)。
