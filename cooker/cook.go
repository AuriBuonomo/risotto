package cooker

import (
	"github.com/auribuo/risotto/log"
	"github.com/spf13/cobra"
)

type riceConfig struct {
	Name        string
	Description string
	Author      string
	Version     string
	Wallpaper   string
	GtkTheme    string
	ConfigFiles []riceConfigFile
}

type riceConfigFile struct {
	FileName   string
	ConfigPath string
}

func Cook(flavor string, restart bool) {
	riceInfo, err := load(flavor)
	cobra.CheckErr(err)
	log.Logf("Cooking %s rice...\n", riceInfo.Name)
	log.Logf("Backing up old config files...\n")
	err = backup(riceInfo)
	cobra.CheckErr(err)
	err = write(riceInfo)
	cobra.CheckErr(err)
	if restart {
		log.Logf("Restarting wm...\n")
		err = restartWm()
		cobra.CheckErr(err)
	}
}

func write(riceInfo riceConfig) error {
	return nil
}
