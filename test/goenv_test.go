package goenv

import (
	"github.com/adeven/goenv"
	"testing"
)

func TestSetEnvironment(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "custom", "")
	if goenv.GetPort() != "6711" {
		t.Error("port != 6711")
	}
}

func TestSetEnvironmentNotFound(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "nonexistent", "")
	if goenv.GetPort() != "8080" {
		t.Error("port != 8080")
	}
}

func TestSetConfigFile(t *testing.T) {
	goenv := goenv.NewGoenv("./config/custom.yml", "config", "")
	if goenv.GetPort() != "4388" {
		t.Error("port != 4388")
	}
}

func TestSetConfigFileNotFound(t *testing.T) {
	defer func() { recover() }()

	goenv.NewGoenv("nonexistent", "", "")
	t.Error("SetConfigFile didn't panic")
}

func TestGet(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "config", "")
	if goenv.Get("custom", "") != "aoeu" {
		t.Error("custom != aoeu")
	}
}

func TestGetNotFound(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "config", "")
	if goenv.Get("nonexistent", "snth") != "snth" {
		t.Error("nonexistent != snth")
	}
}

// TODO: test new functions
