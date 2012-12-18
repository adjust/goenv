package goenv

import (
	"github.com/adeven/go-gypsy/yaml"
	"log"
	"os"
)

var (
	config      *yaml.File
	environment string
)

func init() {
	environment = getEnv("GO_ENV", "development")
	configFile := getEnv("GO_CONFIG", "./config/config.yaml")

	config = yaml.ConfigFile(configFile)
	logFile := Get("log_file", "./log/server.log")
	setLogFile(logFile)

	exitHandler = &StandardHandler{}
	startSignalCatcher()
}

// get value from current environment
func Get(spec, defaultValue string) string {
	value, err := config.Get(environment + "." + spec)
	if err != nil {
		value = defaultValue
	}
	return value
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}

func setLogFile(fileName string) {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Failed to open logFile: " + fileName)
	}
	log.SetOutput(logFile)
	log.SetFlags(5)
}
