package goenv

import (
	"github.com/adeven/goenv"
	"testing"
)

func TestGetPostgres(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("postgres")
	if goenv.GetPostgres() != "user=ter dbname=41 sslmode=disable host=hor" {
		t.Error("postgres != user=ter dbname=41 sslmode=disable host=hor")
	}
}

func TestGetPostgresNotFound(t *testing.T) {
	defer func() { recover() }()

	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("nonexistent")
	goenv.GetPostgres()
	t.Error("GetPostgres didn't panic")
}

func TestGetNamedPostgres(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("postgres")
	if goenv.GetNamedPostgres("custom") != "user=orr dbname=11 sslmode=disable host=obk" {
		t.Error("custom != user=orr dbname=11 sslmode=disable host=obk")
	}
}

func TestGetNamedPostgresNotFound(t *testing.T) {
	defer func() { recover() }()

	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("postgres")
	goenv.GetNamedPostgres("non")
	t.Error("GetNamedPostgres didn't panic")
}
