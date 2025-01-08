package helpers

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./internal/config/config-docker"
	}

	return "./internal/config/config-local"
}
