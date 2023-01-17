package console

import (
	"GoChat/pkg/console"
	"errors"

	"github.com/spf13/cobra"
)

var CmdConsoleName = &cobra.Command{
	Use:   "console_name",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runConsoleName,
}

func runConsoleName(cmd *cobra.Command, args []string) {

	console.Success("this is a console")
	console.Success("这是一条成功的提示")
	console.Warning("这是一条提示")
	console.Error("这是一条错误信息")
	console.Warning("终端输出最好使用英文，这样兼容性会更好~")
	console.Exit("exit 方法可以用来打印消息并中断程序！")
	console.ExitIf(errors.New("在 err = nil 的时候打印并退出"))
}
