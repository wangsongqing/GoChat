package job

import (
	"GoChat/pkg/console"
	"errors"

	"github.com/spf13/cobra"
)

// CmdJobName 运行命令  go run main.go job job_name
var CmdJobName = &cobra.Command{
	Use:   "job_name",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runJobName,
}

func runJobName(cmd *cobra.Command, args []string) {

	console.Success("this is a job")
	console.Success("这是一条成功的提示")
	console.Warning("这是一条提示")
	console.Error("这是一条错误信息")
	console.Warning("终端输出最好使用英文，这样兼容性会更好~")
	console.Exit("exit 方法可以用来打印消息并中断程序！")
	console.ExitIf(errors.New("在 err = nil 的时候打印并退出"))
}
