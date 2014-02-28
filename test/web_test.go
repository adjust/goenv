package goenv

import (
	"github.com/adjust/goenv"
	"testing"
)

func TestGetPort(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "web", "nil")
	if goenv.GetPort() != "3367" {
		t.Error("port != 3367")
	}
}

func TestGetPortNotFound(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "nonexistent", "nil")
	if goenv.GetPort() != "8080" {
		t.Error("port != 8080")
	}
}
