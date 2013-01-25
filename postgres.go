package goenv

import (
	"fmt"
	"log"
)

func GetPostgres() string {
	result := GetNamedPostgres("postgres")
	return result
}

func GetNamedPostgres(name string) string {
	user := Get(name+".user", "postgres")
	host := Get(name+".host", "localhost")
	dbst := Get(name+".db", "")

	if dbst == "" {
		log.Panicf("Missing value in config.yml: %s.%s.db", environment, name)
	}

	result := fmt.Sprintf("user=%s dbname=%s sslmode=disable host=%s", user, dbst, host)
	return result
}
