package goenv

import (
	"log"
	"strconv"
)

func GetPort() string {
	port := Get("port", "8080")
	return port
}

func GetCookieDomain() string {
	domain := Get("cookie_domain", "localhost")
	return domain
}

func GetShard() int64 {
	shardst := Get("shard", "1")
	shard, err := strconv.ParseInt(shardst, 10, 64)
	if err != nil {
		log.Println("Failed to convert to int64: " + shardst)
		shard = 1
	}
	return shard
}
