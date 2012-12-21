package goenv

import (
	"github.com/adeven/goenv"
	"testing"
)

// TODO: add tests that set environment variables instead of calling the setters

func TestSetEnvironment(t *testing.T) {
	goenv.SetConfigFile("./config/config.yml")
	goenv.SetEnvironment("custom")
	if goenv.GetPort() != "6711" {
		t.Error("port != 6711")
	}
}

func TestSetEnvironmentNotFound(t *testing.T) {
	goenv.SetConfigFile("./config/config.yml")
	goenv.SetEnvironment("nonexistent")
	if goenv.GetPort() != "8080" {
		t.Error("port != 8080")
	}
}

func TestSetConfigFile(t *testing.T) {
	goenv.SetConfigFile("./config/custom.yml")
	goenv.SetEnvironment("config")
	if goenv.GetPort() != "4388" {
		t.Error("port != 4388")
	}
}

func TestSetConfigFileNotFound(t *testing.T) {
	defer func() { recover() }()

	goenv.SetConfigFile("nonexistent")
	t.Error("SetConfigFile didn't panic")
}

func TestGet(t *testing.T) {
	goenv.SetConfigFile("./config/config.yml")
	goenv.SetEnvironment("config")
	if goenv.Get("custom", "") != "aoeu" {
		t.Error("custom != aoeu")
	}
}

func TestGetNotFound(t *testing.T) {
	goenv.SetConfigFile("./config/config.yml")
	goenv.SetEnvironment("config")
	if goenv.Get("nonexistent", "snth") != "snth" {
		t.Error("nonexistent != snth")
	}
}
