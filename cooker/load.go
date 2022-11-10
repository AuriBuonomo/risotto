package cooker

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func load(flavour string) (riceConfig, error) {
	templateDir := viper.GetString("templates.templates_path")
	ignorePrefix := viper.GetString("templates.ignore_prefix")
	direcotries, err := os.ReadDir(templateDir)
	cobra.CheckErr(err)
	for i := 0; i < len(direcotries); i++ {
		dir := direcotries[i]
		dirName := dir.Name()

		if !dir.IsDir() {
			continue
		}
		if len(dirName) > len(ignorePrefix) && dirName[0:len(ignorePrefix)] == ignorePrefix {
			continue
		}
		if dirName == flavour {
			configFileName := templateDir + "/" + dirName + "/rice.json"
			configFile, err := os.ReadFile(configFileName)
			if err != nil {
				return riceConfig{}, err
			}
			var riceInfo riceConfig
			err = json.Unmarshal(configFile, &riceInfo)
			if err != nil {
				return riceConfig{}, err
			}
			return riceInfo, nil
		}
	}
	return riceConfig{}, fmt.Errorf("could not find rice with name %s", flavour)
}
