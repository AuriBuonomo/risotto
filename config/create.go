package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var config = fmt.Sprintf(`
# Template config
[templates]

# Template location
# Each template corresponds to a rice
# The directory name is the name of the rice
# The directory must contain a file named "rice.json"
# The template.conf file contains the rice configuration such as the name, description, window manager, etc.
templates_path = "%s/.local/share/risotto"

# Prefix for ignored templates
ignore_prefix = "_"

[risotto]
active_rice = "none"

# The command to run to reload the full rice
# You would put here a command such as "i3-msg restart" to reload the window manager
# Reloading the bar, rofi etc would all happen in the wm manager itself or a script
[risotto.restart]
command = "$HOME/.config/bspwm/bspwmrc"
args = []
`, os.Getenv("HOME"))

func createConfig() error {
	err := os.MkdirAll(configFolder, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(config)

	if err != nil {
		return err
	}

	fmt.Printf("Created config file %s\n", configPath)
	err = viper.ReadInConfig()
	return err
}
