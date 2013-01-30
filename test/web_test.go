package goenv

import (
	"github.com/adeven/goenv"
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "web", "")
	if goenv.GetPort() != "3367" {
		t.Error("port != 3367")
	}
}

func TestGetPortNotFound(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "nonexistent", "")
	if goenv.GetPort() != "8080" {
		t.Error("port != 8080")
	}
}

func TestGetCookieDomain(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "web", "")
	if goenv.GetCookieDomain() != "dadadomain" {
		t.Error("cookie_domain != dadadomain")
	}
}

func TestGetCookieDomainNotFound(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "nonexistent", "")
	if goenv.GetCookieDomain() != "localhost" {
		t.Error("cookie_domain != localhost")
	}
}

func TestGetShard(t *testing.T) {
	os.Setenv("GO_SHARD", "17")
	goenv := goenv.DefaultGoenv()
	if goenv.GetShard() != 17 {
		t.Error("shard != 17")
	}
}

func TestGetShardNotFound(t *testing.T) {
	defer func() {
		recover()
	}()

	os.Setenv("GO_SHARD", "")
	goenv := goenv.DefaultGoenv()
	goenv.GetShard()
	t.Error("GetShard didn't panic")
}
