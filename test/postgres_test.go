package goenv

import (
	"github.com/adeven/goenv"
	"testing"
)

func TestGetPostgres(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "postgres", "nil")
	if goenv.GetPostgres() != "user=ter dbname=41 sslmode=disable host=hor port=4711" {
		t.Error("postgres != user=ter dbname=41 sslmode=disable host=hor port=4711")
	}
}

func TestGetPostgresNotFound(t *testing.T) {
	defer func() { recover() }()

	goenv := goenv.NewGoenv("./config/config.yml", "nonexistent", "nil")
	goenv.GetPostgres()
	t.Error("GetPostgres didn't panic")
}

func TestGetNamedPostgres(t *testing.T) {
	goenv := goenv.NewGoenv("./config/config.yml", "postgres", "nil")
	if goenv.GetNamedPostgres("custom") != "user=orr dbname=11 sslmode=disable host=obk port=5432" {
		t.Error("custom != user=orr dbname=11 sslmode=disable host=obk port=5432")
	}
}

func TestGetNamedPostgresNotFound(t *testing.T) {
	defer func() { recover() }()

	goenv := goenv.NewGoenv("./config/config.yml", "postgres", "nil")
	goenv.GetNamedPostgres("non")
	t.Error("GetNamedPostgres didn't panic")
}
