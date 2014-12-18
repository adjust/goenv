package goenv

import (
	"bytes"
	"github.com/adjust/go-gypsy/yaml"
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

	if logFile == "" {
		logFile = goenv.Get("log_file", "./log/server.log")
		bufferedLog := goenv.GetInt("buffered_log", 0)
		os.MkdirAll(path.Dir(logFile), 0755)
		setLogFile(logFile, bufferedLog)
	}

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

func setLogFile(fileName string, bufferedLog int) {
	if fileName == "nil" {
		return
	}

	var buffer bytes.Buffer
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("goenv failed to open logFile: " + fileName)
	}

	if bufferedLog == 1 {
		log.SetOutput(&buffer)
		go flushBufferToFile(&buffer, logFile)
	} else {
		log.SetOutput(logFile)
	}

	log.SetFlags(5)
}

func flushBufferToFile(buffer *bytes.Buffer, logFile *os.File) {
	buffer.Grow(10 * 1024 * 1024)
	for {
		time.Sleep(500 * time.Millisecond)
		buffer.WriteTo(logFile)
	}
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
