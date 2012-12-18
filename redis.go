package goenv

import (
	"log"
	"strconv"
)

func GetRedis() (host, port string, db int) {
	return GetNamedRedis("redis")
}

func GetNamedRedis(name string) (host, port string, db int) {
	host = Get(name+".host", "localhost")
	port = Get(name+".port", "6379")
	dbs := Get(name+".db", "0")

	db, err := strconv.Atoi(dbs)
	if err != nil {
		log.Panic(name + ": Failed to convert to int: " + dbs)
	}
	return
}
