package goenv

import (
	"log"
	"os"
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
	shardst := Get("shard", "0")
	shard, err := strconv.ParseInt(shardst, 10, 64)
	if err != nil || shard == 0 {
		log.Panic("Invalid shard: \"" + shardst + "\"")
		os.Exit(1)
	}
	return shard
}
