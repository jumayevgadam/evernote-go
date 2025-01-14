package helpers

import "os"

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return os.Getenv("DOCKER_CONFIG_PATH")
	}

	return os.Getenv("LOCAL_CONFIG_PATH")
}
