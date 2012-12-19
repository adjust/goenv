package goenv

import (
	"github.com/adeven/goenv"
	"testing"
)

func TestGetPort(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("web")
	if goenv.GetPort() != "3367" {
		t.Error("port != 3367")
	}
}

func TestGetPortNotFound(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("nonexistent")
	if goenv.GetPort() != "8080" {
		t.Error("port != 8080")
	}
}

func TestGetCookieDomain(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("web")
	if goenv.GetCookieDomain() != "dadadomain" {
		t.Error("cookie_domain != dadadomain")
	}
}

func TestGetCookieDomainNotFound(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("nonexistent")
	if goenv.GetCookieDomain() != "localhost" {
		t.Error("cookie_domain != localhost")
	}
}

func TestGetShard(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("web")
	if goenv.GetShard() != 17 {
		t.Error("shard != 17")
	}
}

func TestGetShardNotFound(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("nonexistent")

	defer func() {
		recover()
	}()

	goenv.GetShard()
	t.Error("GetShard didn't panic")
}
