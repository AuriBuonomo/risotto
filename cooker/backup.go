package cooker

import "os"

func backup(riceInfo riceConfig) error {
	for i := 0; i < len(riceInfo.ConfigFiles); i++ {
		configFile := riceInfo.ConfigFiles[i]
		if _, err := os.Stat(configFile.ConfigPath); err == nil {
			err = os.Remove(configFile.ConfigPath + ".bak")
			if err != nil {
				return err
			}
		}
		err := os.Rename(configFile.ConfigPath, configFile.ConfigPath+".bak")
		if err != nil {
			return err
		}
	}
	return nil
}
