package goenv

import (
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
		log.Panic("Missing value in config.yml: " + environment + "." + name + ".db")
	}

	result := "user=" + user + " dbname=" + dbst + " sslmode=disable host=" + host
	return result
}
