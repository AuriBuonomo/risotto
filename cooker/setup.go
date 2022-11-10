package cooker

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Setup() error {
	templateDir := viper.GetString("templates.templates_path")
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		err = os.Mkdir(templateDir, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("Created template directory at %s\n", templateDir)
	}
	return nil
}
