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
	shardst := getEnv("GO_SHARD", "")
	shard, err := strconv.ParseInt(shardst, 10, 64)
	if err != nil || shard == 0 {
		log.Panic("invalid shard: \"" + shardst + "\"")
	}
	return shard
}
