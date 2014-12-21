package goenv

import (
	"testing"
)

func TestGetGraphite(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "graphite", "nil")
	host, port := goenv.GetGraphite()
	if host != "ghost" || port != 4177 {
		t.Error("graphite != ghost, 4177")
	}
}

func TestGetGraphiteNotFound(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "nonexistent", "nil")
	host, port := goenv.GetGraphite()
	if host != "" || port != 2003 {
		t.Error("graphite != '', 2003")
	}
}
