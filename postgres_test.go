package goenv

import (
	"testing"
)

func TestGetPostgres(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "postgres", "nil")
	if goenv.GetPostgres() != "user=ter dbname=41 sslmode=disable host=hor port=4711" {
		t.Error("postgres != user=ter dbname=41 sslmode=disable host=hor port=4711")
	}
}

func TestGetPostgresNotFound(t *testing.T) {
	defer func() { recover() }()

	goenv := NewGoenv("./test_config.yml", "nonexistent", "nil")
	goenv.GetPostgres()
	t.Error("GetPostgres didn't panic")
}

func TestGetNamedPostgres(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "postgres", "nil")
	if goenv.GetNamedPostgres("custom") != "user=orr dbname=11 sslmode=disable host=obk port=5432" {
		t.Error("custom != user=orr dbname=11 sslmode=disable host=obk port=5432")
	}
}

func TestGetNamedPostgresNotFound(t *testing.T) {
	defer func() { recover() }()

	goenv := NewGoenv("./test_config.yml", "postgres", "nil")
	goenv.GetNamedPostgres("non")
	t.Error("GetNamedPostgres didn't panic")
}
