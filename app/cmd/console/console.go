package console

import (
	"github.com/spf13/cobra"
)

// Console Job 说明 cobra 命令
var Console = &cobra.Command{
	Use:   "console",
	Short: "Generate console file and code",
}

func init() {
	Console.AddCommand(
		CmdConsoleName,
	)
}
