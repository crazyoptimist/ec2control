package utils

import (
	"errors"
	"os"
	"runtime"
)

func GetConfigDir() (string, error) {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		configDir := os.ExpandEnv("$HOME/.config")

		// Create it if not existing
		_ = os.Mkdir(configDir, 0700)
		return configDir, nil
	} else if runtime.GOOS == "windows" {
		roamingDir, _ := os.UserConfigDir()
		configDir := roamingDir + "\\ec2control"

		_ = os.Mkdir(configDir, 0700)
		return configDir, nil
	}
	return "", errors.New("Cound not find the config dir: Unsupported OS")
}
