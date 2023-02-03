package job

import (
	v1 "GoChat/app/http/controllers/api/v1"
	"GoChat/app/models/chat_log"
	"GoChat/pkg/redis"
	"encoding/json"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"time"
)

// CmdPopChatLog 运行命令 go run main.go job pop_chat_log
var CmdPopChatLog = &cobra.Command{
	Use:   "pop_chat_log",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runPopChatLog,
}

func runPopChatLog(cmd *cobra.Command, args []string) {

	key := "go_chat_log"

	for {
		data := redis.Redis.Rpop(key)
		if data == "" {
			time.Sleep(time.Second)
			color.Cyanln("暂无聊天记录")
			continue
		}

		message := v1.Message{}
		json.Unmarshal([]byte(data), &message)

		chatLog := chat_log.ChatLog{}
		chatLog.UserId = int(message.FormId)
		chatLog.TargetId = int(message.TargetId)
		chatLog.Type = message.Type
		chatLog.Media = message.Media
		chatLog.Content = message.Content
		if ok := chatLog.Create(); ok == 0 {
			zap.S().Info("持久化失败", message)
		}

	}

}
