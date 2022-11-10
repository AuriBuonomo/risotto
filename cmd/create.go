package cmd

import (
	"github.com/auribuo/risotto/cooker"
	"github.com/spf13/cobra"
)

var createCommand = &cobra.Command{
	Use:   "create [flags] [name]",
	Short: "create a new rice flavour",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cooker.CreateFlavour(args[0])
	},
}

func init() {
	flavourCommand.AddCommand(createCommand)
}
