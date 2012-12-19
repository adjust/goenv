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
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("nonexistent")
	if goenv.GetPostgres() != "user=postgres dbname=0 sslmode=disable host=localhost" {
		t.Error("postgres != user=postgres dbname=0 sslmode=disable host=localhost")
	}
}

func TestGetNamedPostgres(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("postgres")
	if goenv.GetNamedPostgres("custom") != "user=orr dbname=11 sslmode=disable host=obk" {
		t.Error("custom != user=orr dbname=11 sslmode=disable host=obk")
	}
}

func TestGetNamedPostgresNotFound(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("postgres")
	if goenv.GetNamedPostgres("non") != "user=postgres dbname=0 sslmode=disable host=localhost" {
		t.Error("nonexistent != user=postgres dbname=0 sslmode=disable host=localhost")
	}
}
