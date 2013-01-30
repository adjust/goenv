package goenv

import (
	"log"
	"strconv"
)

func (goenv *Goenv) GetPort() string {
	port := goenv.Get("port", "8080")
	return port
}

func (goenv *Goenv) GetCookieDomain() string {
	domain := goenv.Get("cookie_domain", "localhost")
	return domain
}

func (goenv *Goenv) GetShard() int64 {
	shardst := getEnv("GO_SHARD", "")
	shard, err := strconv.ParseInt(shardst, 10, 64)
	if err != nil || shard == 0 {
		log.Panic("invalid shard: \"" + shardst + "\"")
	}
	return shard
}
