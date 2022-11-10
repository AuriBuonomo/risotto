package main

import (
	"github.com/auribuo/risotto/cmd"
	"github.com/auribuo/risotto/config"
	"github.com/auribuo/risotto/cooker"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(config.Setup())
	cobra.CheckErr(cooker.Setup())
	cmd.Execute()
}
