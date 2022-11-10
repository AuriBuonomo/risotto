package cmd

import "github.com/spf13/cobra"

var recipeCommand = &cobra.Command{
	Use:   "recipe",
	Short: "manage your risotto config",
	Long: `
recipe is used to alter the config file for risotto
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() //TODO: implement
	},
}

func init() {
	rootCommand.AddCommand(recipeCommand)
}
