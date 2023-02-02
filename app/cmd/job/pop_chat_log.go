package job

import (
	v1 "GoChat/app/http/controllers/api/v1"
	"GoChat/pkg/redis"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

// CmdPopChatLog 运行命令 go run main.go job pop_chat_log
var CmdPopChatLog = &cobra.Command{
	Use:   "pop_chat_log",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runPopChatLog,
}

func runPopChatLog(cmd *cobra.Command, args []string) {

	key := "go_chat_log"
	data := redis.Redis.Rpop(key)

	message := v1.Message{}
	json.Unmarshal([]byte(data), &message)
	fmt.Println(message.TargetId)
	fmt.Println(message.Type)
}
