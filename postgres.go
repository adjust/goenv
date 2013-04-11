package goenv

import (
	"fmt"
	"log"
)

func (goenv *Goenv) GetPostgres() string {
	result := goenv.GetNamedPostgres("postgres")
	return result
}

func (goenv *Goenv) GetNamedPostgres(name string) string {
	user := goenv.Get(name+".user", "postgres")
	host := goenv.Get(name+".host", "localhost")
	port := goenv.Get(name+".port", "5432")
	dbst := goenv.Get(name+".db", "")

	if dbst == "" {
		log.Panicf("Missing value in config.yml: %s.%s.db", goenv.environment, name)
	}

	result := fmt.Sprintf("user=%s dbname=%s sslmode=disable host=%s port=%s", user, dbst, host, port)
	return result
}
