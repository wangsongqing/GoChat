package make

import (
	"GoChat/pkg/console"
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeJOB = &cobra.Command{
	Use:   "job",
	Short: "Create a command, should be snake_case, exmaple: make cmd buckup_database",
	Run:   runMakeJOB,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeJOB(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/cmd/job/%s.go", model.PackageName)

	// 从模板中创建文件（做好变量替换）
	createFileFromStub(filePath, "job", model)

	// 友好提示
	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
