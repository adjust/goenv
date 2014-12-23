package goenv

import (
	"testing"
)

func TestGetPort(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "web", "nil")
	if goenv.GetPort() != "3367" {
		t.Error("port != 3367")
	}
}

func TestGetPortNotFound(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "nonexistent", "nil")
	if goenv.GetPort() != "8080" {
		t.Error("port != 8080")
	}
}
