package cooker

import (
	"os/exec"

	"github.com/spf13/viper"
)

func restartWm() error {
	command := viper.GetString("risotto.restart.command")
	args := viper.GetStringSlice("risotto.restart.args")
	err := exec.Command(command, args...).Run()
	return err
}
