package cmd

import (
	"github.com/auribuo/risotto/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var doLogs = false

var rootCommand = &cobra.Command{
	Use:   "risotto",
	Short: "multi rice management tool",
	Long: `
risotto is a tool for managing your rice setups
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.SetVerbose(doLogs)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	Version: "1.0.0",
}

func init() {
	rootCommand.PersistentFlags().BoolVarP(&doLogs, "log", "l", false, "output text logs to stdout")

	rootCommand.SetVersionTemplate("{{.Version}}\n")
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:       rootCommand,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		Commands:      cc.HiYellow + cc.Bold,
		Example:       cc.Italic,
		ExecName:      cc.Bold,
		Flags:         cc.Bold,
		FlagsDataType: cc.Italic + cc.HiBlue,
	})

	rootCommand.Execute()
}
