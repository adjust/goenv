package goenv

import (
	"github.com/adeven/goenv"
	"testing"
)

func TestGetRedis(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "redis", "")
	host, port, db := goenv.GetRedis()
	if host != "ecr" || port != "340" || db != 16 {
		t.Error("redis != ecr, 340, 16")
	}
}

func TestGetRedisNotFound(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "nonexistent", "")
	host, port, db := goenv.GetRedis()
	if host != "localhost" || port != "6379" || db != 0 {
		t.Error("redis != localhost, 6379, 0")
	}
}

func TestGetNamedRedis(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "redis", "")
	host, port, db := goenv.GetNamedRedis("custom")
	if host != "ruo" || port != "114" || db != 81 {
		t.Error("custom != ruo, 114, 81")
	}
}

func TestGetNamedRedisNotFound(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "redis", "")
	host, port, db := goenv.GetNamedRedis("nonexistent")
	if host != "localhost" || port != "6379" || db != 0 {
		t.Error("nonexistent != localhost, 6379, 0")
	}
}
