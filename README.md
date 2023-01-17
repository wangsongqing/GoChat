# GoChat
## Golang开发的一个简单的聊天系统

### 用到的组件
- MySQL -- 数据库操作
- Redis -- redis操作
- zap -- 高新能日志方案
- Gorm -- 数据库操作
- Cobra -- 命令行结构
- limiter -- 限流器
- Jwt -- JWT操作
- 构建docker镜像

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
- 运行项目
```azure
go run main.go
```

[测试链接](http://www.websocket-test.com/)

链接websocket(token为登录用户的token)
- 用户1:ws://127.0.0.1:8023/v1/msg/send?token=xxx
- 用户2:ws://127.0.0.1:8023/v1/msg/send?token=xxx

发送消息的格式
```json
{
  "TargetId":1, // 接受对象的ID
  "Type":1, // 消息的类型
  "CreateTime":1672996855236,
  "userId":2, // 发送对象的ID
  "Media":1,
  "Content":"在干嘛" // 消息的内容
}
```
---

框架开发方式具体请看:[GoLaravel框架](https://github.com/wangsongqing/GoLaravel)


