# 小程序 code2Session 示例

用 `wx.login()` 拿到的 `code` 换取用户身份，结果直接打印到控制台。

## 运行

```bash
export WECHAT_APPID=wx1234567890abcdef
export WECHAT_APPSECRET=0123456789abcdef0123456789abcdef

go run ./miniProgram/code2Session -code=0d3d1NFa1YlqKM0koXFa1yGSd22d1NFt
```

输出示例：

```
2026/11/07 20:00:47 openid: oKGug4tqqGYFCB46wZLoMBC5XN9Q
2026/11/07 20:00:47 unionid: uKGug4tqqGYFCB46wZLoMBC5XN9Q
2026/11/07 20:00:47 session_key: iriBgaEJgBlAj4PXBP3txA==
```

未绑定开放平台时 `unionid` 为空。

## 说明

- `code` 一次性有效，5 分钟过期，失败后需重新调用 `wx.login()`；
- `session_key` 是解密手机号等数据的密钥，只应保存在服务端，不要返回给客户端或写入日志；
- 微信正常回包但拒绝请求时（如 40029 code 无效、40226 高风险用户、45011 调用太频繁），
  示例通过 `res.OK()` 判断并打印可读错误。
