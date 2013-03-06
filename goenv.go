package goenv

import (
	"github.com/adeven/go-gypsy/yaml"
	"log"
	"os"
	"strconv"
)

type Goenv struct {
	configFile  *yaml.File
	environment string
}

func NewGoenv(configFile, environment, logFile string) *Goenv {
	if environment == "" {
		environment = "development"
	}

	goenv := &Goenv{
		configFile:  yaml.ConfigFile(configFile),
		environment: environment,
	}

	if goenv.configFile == nil {
		panic("goenv failed to open configFile: " + configFile)
	}

	if logFile == "" {
		logFile = goenv.Get("log_file", "./log/server.log")
	}
	setLogFile(logFile)

	return goenv
}

func DefaultGoenv() *Goenv {
	environment := getEnv("GO_ENV", "development")
	configFile := getEnv("GO_CONFIG", "./config.yml")
	return NewGoenv(configFile, environment, "")
}

func setLogFile(fileName string) {
	if fileName == "nil" {
		return
	}

	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("goenv failed to open logFile: " + fileName)
	}
	log.SetOutput(logFile)
	log.SetFlags(5)
}

// get value from current environment
func (goenv *Goenv) Get(spec, defaultValue string) string {
	value, err := goenv.configFile.Get(goenv.environment + "." + spec)
	if err != nil {
		value = defaultValue
	}
	return value
}

func (goenv *Goenv) GetInt(spec string, defaultValue int) int {
	str := goenv.Get(spec, "")
	if str == "" {
		return defaultValue
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("goenv GetInt failed Atoi", goenv.environment, spec, str)
	}
	return val
}

func (goenv *Goenv) Require(spec string) string {
	value := goenv.Get(spec, "")
	if value == "" {
		log.Panicf("goenv Require couldn't find %s.%s", goenv.environment, spec)
	}
	return value
}

func (goenv *Goenv) RequireInt(spec string) int {
	str := goenv.Require(spec)
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("goenv RequireInt failed Atoi", goenv.environment, spec, str)
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
