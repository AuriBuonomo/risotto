package cooker

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const riceConfigSchema = `
{
    "type": "object",
    "properties": {
        "name": {
            "type": "string"
        },
        "description": {
            "type": "string"
        },
        "author": {
            "type": "string"
        },
        "version": {
            "type": "string",
            "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$"
        },
        "wallpaper": {
            "type": "string",
            "format": "uri"
        },
        "gtkTheme": {
            "type": "string"
        },
        "configFiles": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "fileName": {
                        "type": "string"
                    },
                    "configPath": {
                        "type": "string",
                        "format": "uri"
                    }
                } 
            }
        }
    }
}
`

const riceConfigTemplate = `
{
	"schema": "./rice.schema.json",
	"name": "{{.Name}}",
	"description": "{{.Description}}",
	"author": "{{.Author}}",
	"version": "{{.Version}}",
	"wallpaper": "{{.Wallpaper}}",
	"gtkTheme": "{{.GtkTheme}}",
	"configFiles": []
}
`

func RemoveFlavour(flavour string) {
	templateDir := viper.GetString("templates.templates_path")
	ignorePrefix := viper.GetString("templates.ignore_prefix")
	direcotries, err := os.ReadDir(templateDir)
	hasRemoved := false
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
			err = os.RemoveAll(templateDir + "/" + dirName)
			cobra.CheckErr(err)
			hasRemoved = true
		}
	}
	if !hasRemoved {
		cobra.CheckErr(fmt.Errorf("could not find rice with name %s", flavour))
	}
}

func ListAllFlavours() {
	fmt.Println("Available flavours:")
	templateDir := viper.GetString("templates.templates_path")
	ignorePrefix := viper.GetString("templates.ignore_prefix")
	direcotries, err := os.ReadDir(templateDir)
	ctr := 1
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
		fmt.Printf("%d - %s (%s)\n", ctr, dirName, templateDir+"/"+dirName)
		ctr++
	}
}

func ListCurrentFlavour() {
	flavour := viper.GetString("risotto.active_rice")
	fmt.Printf("Current flavour: %s\n", flavour)
}

func CreateFlavour(flavour string) {
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
			cobra.CheckErr(fmt.Errorf("flavour %s already exists", flavour))
		}
	}
	err = os.Mkdir(templateDir+"/"+flavour, os.ModePerm)
	cobra.CheckErr(err)
	err = os.WriteFile(templateDir+"/"+flavour+"/rice.schema.json", []byte(riceConfigSchema), os.ModePerm)
	cobra.CheckErr(err)

	currentUser, err := user.Current()
	cobra.CheckErr(err)

	riceInfo := riceConfig{
		Name:        flavour,
		Description: "A rice flavour created by risotto",
		Version:     "0.0.1",
		Author:      currentUser.Username,
		Wallpaper:   "wallpaper.png",
		GtkTheme:    "Adwaita",
		ConfigFiles: []riceConfigFile{},
	}

	template, err := template.New("riceConfig").Parse(riceConfigTemplate)
	cobra.CheckErr(err)

	stringWriter := &strings.Builder{}

	err = template.Execute(stringWriter, riceInfo)
	cobra.CheckErr(err)

	err = os.WriteFile(templateDir+"/"+flavour+"/rice.json", []byte(stringWriter.String()), os.ModePerm)
	cobra.CheckErr(err)
	fmt.Printf("Created flavour: %s (%s)\n", flavour, templateDir+"/"+flavour)
}

func EditFlavour(flavour string) {
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
			cmd := exec.Command(os.Getenv("EDITOR"), templateDir+"/"+dirName+"/rice.json")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			cobra.CheckErr(err)
			return
		}
	}
	cobra.CheckErr(fmt.Errorf("could not find rice with name %s", flavour))
}

func AddConfigFileToFlavour(flavour string, configFileName string, configFilePath string) {
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
			riceInfo := riceConfig{}
			riceConfigFilePath := templateDir + "/" + dirName + "/rice.json"
			configContent, err := os.ReadFile(riceConfigFilePath)
			cobra.CheckErr(err)
			err = json.Unmarshal([]byte(configContent), &riceInfo)
			cobra.CheckErr(err)

			riceConfigFileEntry := riceConfigFile{
				FileName:   configFileName,
				ConfigPath: configFilePath,
			}

			riceInfo.ConfigFiles = append(riceInfo.ConfigFiles, riceConfigFileEntry)
			riceConfigFileBytes, err := json.MarshalIndent(riceInfo, "", "    ")
			cobra.CheckErr(err)
			err = os.WriteFile(templateDir+"/"+dirName+"/rice.json", riceConfigFileBytes, os.ModePerm)
			cobra.CheckErr(err)
			return
		}
	}
	cobra.CheckErr(fmt.Errorf("could not find rice with name %s", flavour))
}
