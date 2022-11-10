package cmd

import (
	"github.com/auribuo/risotto/cooker"
	"github.com/spf13/cobra"
)

var cookCommand = &cobra.Command{
	Use:   "cook [flags] [name]",
	Short: "switch between rice setups",
	Long: `
cook is used to switch between rice setups. The command will generate a new config file for the rice setup, backup the old ones and restart the wm
if the flag is set
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		doRestart, err := cmd.Flags().GetBool("restart")
		cobra.CheckErr(err)
		cooker.Cook(args[0], doRestart)
	},
}

func init() {
	cookCommand.Flags().BoolP("restart", "r", false, "restart the wm after switching")

	rootCommand.AddCommand(cookCommand)
}
