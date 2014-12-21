package goenv

import (
	"github.com/adjust/redis"

	"fmt"
)

type RedisLogWriter struct {
	redisClient *redis.Client
	logName     string
}

func (goenv *Goenv) NewRedisLogWriter(logName string, poolSize int) *RedisLogWriter {
	host, port, db := goenv.GetNamedRedis("redis_log")
	network := "tcp"
	addr := fmt.Sprintf("%s:%s", host, port)

	if port == "" {
		network = "unix"
		addr = host
	}

	logWriter := &RedisLogWriter{
		logName: logName,
	}

	options := &redis.Options{
		Network: network,
		Addr:    addr,
		DB:      int64(db),

		PoolSize: poolSize,
	}

	logWriter.redisClient = redis.NewClient(options)
	return logWriter
}

func (logWriter *RedisLogWriter) Write(p []byte) (n int, err error) {
	err = logWriter.redisClient.LPush(logWriter.logName, string(p)).Err()
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
