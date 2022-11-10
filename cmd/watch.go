package cmd

import "github.com/spf13/cobra"

var watchCommand = &cobra.Command{
	Use:   "watch",
	Short: "watch your rice setup",
	Long: `
watch watches the config templates for the current rice setup and updates the wm on change
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() //TODO: implement
	},
}

func init() {
	rootCommand.AddCommand(watchCommand)
}
