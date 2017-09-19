package goenv

import (
	"os"
	"testing"
)

const (
	REDIS_HOST = "my_redis_host"
	REDIS_PORT = "1234"
	ENV = "development"
)

func init() {
	os.Setenv("REDIS_PORT", REDIS_PORT)
	os.Setenv("REDIS_HOST", REDIS_HOST)
	os.Setenv("DEV_ENV", ENV)
}

func TestNewTemplateGoenv(t *testing.T) {
	goenv := NewTemplateGoenv("./test_template.yml", "", "nil")
	host, port, _ := goenv.GetRedis()
	if host != REDIS_HOST {
		t.Errorf("host != %s, got %s", REDIS_HOST, host)
	}

	if port != REDIS_PORT {
		t.Errorf("port != %s, got %s", REDIS_PORT, port)
	}
}
