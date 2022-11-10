package cmd

import (
	"github.com/auribuo/risotto/cooker"
	"github.com/spf13/cobra"
)

var editCommand = &cobra.Command{
	Use:   "edit [flags] [name]",
	Short: "edit a rice flavour",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addFile, err := cmd.Flags().GetBool("add-file")
		cobra.CheckErr(err)
		if addFile {
			fileName, err := cmd.Flags().GetString("file-name")
			cobra.CheckErr(err)
			filePath, err := cmd.Flags().GetString("file-path")
			cobra.CheckErr(err)
			cooker.AddConfigFileToFlavour(args[0], fileName, filePath)
		} else {
			cooker.EditFlavour(args[0])
		}
	},
}

func init() {
	editCommand.Flags().Bool("add-file", false, "add a file to the flavour config")
	editCommand.Flags().StringP("file-name", "n", "", "name of the file to add")
	editCommand.Flags().StringP("file-path", "p", "", "path of the file to add")

	flavourCommand.AddCommand(editCommand)
}
