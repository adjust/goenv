package goenv

import (
	"github.com/adeven/go-gypsy/yaml"
	"log"
	"os"
	"strconv"
)

var (
	config      *yaml.File
	environment string
)

func init() {
	SetEnvironment(getEnv("GO_ENV", "development"))
	SetConfigFile(getEnv("GO_CONFIG", "./config/config.yml"))
	SetLogFile(Get("log_file", "./log/server.log"))

	exitHandler = &StandardHandler{}
	startSignalCatcher()
}

func SetEnvironment(env string) {
	environment = env
}

func SetConfigFile(fileName string) {
	config = yaml.ConfigFile(fileName)
}

func SetLogFile(fileName string) {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Failed to open logFile: " + fileName)
	}
	log.SetOutput(logFile)
	log.SetFlags(5)
}

// get value from current environment
func Get(spec, defaultValue string) string {
	value, err := config.Get(environment + "." + spec)
	if err != nil {
		value = defaultValue
	}
	return value
}

func GetInt(spec string, defaultValue int) int {
	str := Get(spec, "")
	if str == "" {
		return defaultValue
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("goenv GetInt failed Atoi", environment, spec, str)
	}
	return val
}

func Require(spec string) string {
	value := Get(spec, "")
	if value == "" {
		log.Panicf("goenv Require couldn't find %s.%s", environment, spec)
	}
	return value
}

func RequireInt(spec string) int {
	str := Require(spec)
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("goenv RequireInt failed Atoi", environment, spec, str)
	}
	return val
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}
