package cmd

import (
	"github.com/auribuo/risotto/cooker"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:   "remove [flags] [name]",
	Short: "remove a rice flavour",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cooker.RemoveFlavour(args[0])
	},
}

func init() {
	flavourCommand.AddCommand(removeCommand)
}
