package env

import (
	"os"
)

func GetAppPort() string {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	return appPort
}
