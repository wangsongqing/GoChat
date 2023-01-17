package job

import (
	"github.com/spf13/cobra"
)

// Job 说明 cobra 命令
var Job = &cobra.Command{
	Use:   "job",
	Short: "Generate job file and code",
}

func init() {
	Job.AddCommand(
		CmdJobName,
	)
}
