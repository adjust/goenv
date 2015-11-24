package goenv

import (
	"github.com/adjust/go-gypsy/yaml"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

type Goenv struct {
	configFile  *yaml.File
	environment string
}

func New(fallbackConfigFile, fallbackEnvironment string) *Goenv {
	configFilePath := GetEnv("GO_CONFIG", fallbackConfigFile)
	configFile := yaml.ConfigFile(configFilePath)
	if configFile == nil {
		panic("goenv failed to open configFile: " + configFilePath)
	}

	environment := GetEnv("GO_ENV", fallbackEnvironment)
	return &Goenv{
		configFile:  configFile,
		environment: environment,
	}
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
		os.MkdirAll(path.Dir(logFile), 0755)
		setLogFile(logFile)
	}

	return goenv
}

func (goenv *Goenv) SetLogger(writer io.Writer) {
	log.SetOutput(writer)
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

func (goenv *Goenv) GetDuration(spec string, defaultValue string) time.Duration {
	str := goenv.Get(spec, "")
	if str == "" {
		str = defaultValue
	}
	duration, err := time.ParseDuration(str)
	if err != nil {
		log.Panic("goenv GetDuration failed ParseDuration", goenv.environment, spec, str)
	}
	return duration
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

func (goenv *Goenv) RequireDuration(spec string) time.Duration {
	str := goenv.Require(spec)
	duration, err := time.ParseDuration(str)
	if err != nil {
		log.Panic("goenv RequireDuration failed ParseDuration", goenv.environment, spec, str)
	}
	return duration
}

func (goenv *Goenv) Count(spec string) int {
	count, err := goenv.configFile.Count(goenv.environment + "." + spec)
	if err != nil {
		log.Panicf("goenv Count failed %s", err)
	}
	return count
}

func (goenv *Goenv) CountOk(spec string) (count int, ok bool) {
	count, err := goenv.configFile.Count(goenv.environment + "." + spec)
	if err != nil {
		return 0, false
	}
	return count, true
}

func (goenv *Goenv) GetEnvName() string {
	return goenv.environment
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}

func DefaultGoenv() *Goenv {
	environment := GetEnv("GO_ENV", "development")
	configFile := GetEnv("GO_CONFIG", "./config.yml")
	return NewGoenv(configFile, environment, "")
}

func TestGoenv() *Goenv {
	environment := GetEnv("GO_ENV", "testing")
	configFile := GetEnv("GO_CONFIG", "../run/config.yml")
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
