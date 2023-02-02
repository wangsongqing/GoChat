package v1

import (
	"GoChat/app/models/group_people"
	"GoChat/pkg/jwt"
	"GoChat/pkg/redis"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
	"net/http"
	"strconv"
	"sync"
)

type MessageController struct {
	BaseAPIController
}

//Node 构造连接
type Node struct {
	Conn      *websocket.Conn //连接
	Addr      string          //客户端地址
	DataQueue chan []byte     //消息
	GroupSets set.Interface   //好友 / 群
}

type Message struct {
	FormId   int64  `json:"userId"`   //信息发送者
	TargetId int64  `json:"targetId"` //信息接收者
	Type     int    //聊天类型：群聊 私聊 广播
	Media    int    //信息类型：文字 图片 音频
	Content  string //消息内容
	Pic      string `json:"pic"` //图片相关
	Url      string //文件相关
	Desc     string //文件描述
	Amount   int    //其他数据大小
}

//映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

//读写锁
var rwLocker sync.RWMutex

func (m *MessageController) SendMsg(c *gin.Context) {
	Chat(c.Writer, c.Request, c)
}

//Chat    需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	//1.  获取参数 并 检验 token 等合法性
	claims, err := jwt.NewJWT().ParserTokenWebSocket(c)
	userId, err := strconv.ParseInt(claims.UserID, 10, 64)
	if err != nil {
		zap.S().Info("类型转换失败", err)
		return
	}

	//升级为socket
	var isvalida = true
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取socket连接,构造消息节点
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//用户关系

	//将userId和Node绑定
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	fmt.Println("uid", userId)

	//发送接收消息
	//发送消息
	go sendProc(node)
	//接收消息
	go recProc(node)

	sendMsg(userId, []byte("欢迎进入聊天系统"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				zap.S().Info("写入消息失败", err)
				return
			}
			fmt.Println("数据发送socket成功")
		}

	}
}

//recProc 从websocket中将消息体拿出，然后进行解析，再进行信息类型判断， 最后将消息发送至目的用户的node中
func recProc(node *Node) {
	for {
		//获取信息
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("读取消息失败", err)
			return
		}

		//这里是简单实现的一种方法
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			zap.S().Info("json解析失败", err)
			return
		}

		if msg.Type == 1 {
			zap.S().Info("这是一条私信:", msg.Content)
			tarNode, ok := clientMap[msg.TargetId]
			if !ok {
				zap.S().Info("不存在对应的node", msg.TargetId)
				return
			}

			tarNode.DataQueue <- data
			fmt.Println("发送成功：", string(data))

			// 记录聊天信息到redis，异步写入数据库
			LogKey := "go_chat_log"
			data, _ := json.Marshal(&msg)
			if ok := redis.Redis.Lpush(LogKey, data); ok == false {
				zap.S().Info("聊天记录写入redis失败", msg.TargetId)
			}
		}

		// 群聊
		if msg.Type == 2 {
			groupPeople := group_people.GetGroupMan(msg.TargetId)
			for _, person := range groupPeople {
				if person.UserId == int(msg.FormId) { // 群消息不推送给自己
					continue
				}
				NodeId := int64(person.UserId)
				tarNode, ok := clientMap[NodeId]
				if !ok {
					zap.S().Info("用户没在线", NodeId)
				}
				tarNode.DataQueue <- data
			}
		}

	}
}

//sendMs 向用户发送消息
func sendMsg(id int64, msg []byte) {
	rwLocker.Lock()
	node, ok := clientMap[id]
	rwLocker.Unlock()

	if !ok {
		zap.S().Info("userID没有对应的node")
		return
	}

	zap.S().Info("targetID:", id, "node:", node)
	if ok {
		node.DataQueue <- msg
	}
}
