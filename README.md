# GoChat
## Golang开发的一个简单的聊天系统

### 实现功能点
- 用户手机号注册
- 用户邮箱注册
- 用户手机登录
- 用户邮箱登录
- 获取用户信息
- 获取好友列表
- 获取聊天记录
- 添加好友
- 好友私聊
- 新建群
- 获取用户群列表
- 添加用户进群
- 用户退群
- 群聊


### 用到的组件
- MySQL -- 数据库操作
- Redis -- redis操作
- zap -- 高新能日志方案
- Gorm -- 数据库操作
- Cobra -- 命令行结构
- Jwt -- JWT操作

### 1. 运行项目

- 拉取并安装
```go
git clone git@github.com:wangsongqing/GoChat.git
go mod tidy
```
- 修改 .env.example 为 .env
```azure
cp .env.example .env 
```

- 新建数据库
```go
create database go_chat
```

- 迁移数据库
```azure
go run main.go migrate up
```

- 配置redis
```json
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_CACHE_DB=0
REDIS_MAIN_DB=1
```
- 运行项目
```azure
go run main.go
```

- 运行job脚本同步聊天记录到到数据
```json
go run main.go job pop_chat_log
```

[测试链接](http://www.websocket-test.com/)

链接websocket(token为登录用户的token)
- 用户1:ws://127.0.0.1:8023/v1/msg/send?token=xxx
- 用户2:ws://127.0.0.1:8023/v1/msg/send?token=xxx

发送消息的格式
```json
{
  "TargetId":1, // 信息接收者ID,如果是群则为群ID
  "Type":1, // 消息的类型: 1私聊  2群聊
  "CreateTime":1675050402,// 消息时间戳
  "userId":2, // 信息发送者ID
  "Media":1, // 信息类型:1文字
  "Content":"在干嘛呢？" // 消息的内容
}
```
---

该项目的开发框架是我自己写的一个开源项目，具体请点击链接:[GoLaravel框架](https://github.com/wangsongqing/GoLaravel)


