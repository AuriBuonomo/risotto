package cmd

import (
	"github.com/spf13/cobra"
)

var flavourCommand = &cobra.Command{
	Use:   "flavour",
	Short: "manage your rice flavours",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCommand.AddCommand(flavourCommand)
}
