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

	goenv.setLogFile(logFile)

	return goenv
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

func (goenv *Goenv) setLogFile(fileName string) {
	if fileName == "nil" {
		return
	}

	if fileName == "" {
		fileName = goenv.Get("log_file", "./log/server.log")
	}

	var logFile io.Writer
	var err error

	useRedisLogWriter := goenv.GetBool("log_to_redis", false)

	if useRedisLogWriter {
		logFile, err = goenv.NewRedisLogWriter(fileName)

	} else {
		logFile, err = goenv.getFileHandler(fileName)
	}

	if err != nil {
		panic("goenv failed to open logFile: " + fileName)
	}

	log.SetOutput(logFile)
	log.SetFlags(5)
}

func (goenv *Goenv) getFileHandler(fileName string) (logFile *os.File, err error) {
	os.MkdirAll(path.Dir(fileName), 0755)
	logFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	return
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

func (goenv *Goenv) GetBool(spec string, defaultValue bool) bool {
	str := goenv.Get(spec, "")
	if str == "" {
		return defaultValue
	}

	if str == "true" {
		return true
	}

	if str == "false" {
		return false
	}

	log.Panic("goenv GetBool failed to read value", goenv.environment, spec, str)
	return false
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
