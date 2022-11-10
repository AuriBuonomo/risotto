package cmd

import (
	"github.com/auribuo/risotto/cooker"
	"github.com/spf13/cobra"
)

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "show the current rice flavour",
	Run: func(cmd *cobra.Command, args []string) {
		if all, _ := cmd.Flags().GetBool("all"); all {
			cooker.ListAllFlavours()
		} else {
			cooker.ListCurrentFlavour()
		}
	},
}

func init() {
	showCommand.Flags().BoolP("all", "a", false, "show all flavours")

	flavourCommand.AddCommand(showCommand)
}
