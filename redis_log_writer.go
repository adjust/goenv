package goenv

import (
	redis "github.com/adjust/redis-latest-head"

	"fmt"
	"strings"
)

type RedisLogWriter struct {
	redisClient  *redis.Client
	logName      string
	inputChannel chan string
}

func (goenv *Goenv) NewRedisLogWriter(logName string) (logWriter *RedisLogWriter, err error) {
	host, port, db := goenv.GetNamedRedis("redis_log_writer")
	poolSize := goenv.GetInt("redis_log_writer.pool_size", 40)
	network := "tcp"
	addr := fmt.Sprintf("%s:%s", host, port)

	if strings.Contains(host, ".sock") {
		network = "unix"
		addr = host
	}

	logWriter = &RedisLogWriter{
		logName: logName,
	}

	options := &redis.Options{
		Network: network,
		Addr:    addr,
		DB:      int64(db),

		PoolSize: poolSize,
	}

	logWriter.redisClient = redis.NewClient(options)
	err = logWriter.redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}
	logWriter.inputChannel = make(chan string, 10000)
	logWriter.startConsumers(poolSize)

	return logWriter, nil
}

func (logWriter *RedisLogWriter) Write(p []byte) (n int, err error) {
	logWriter.inputChannel <- string(p)
	return len(p), nil
}

func (logWriter *RedisLogWriter) startConsumers(poolSize int) {
	for i := 0; i < poolSize; i++ {
		go logWriter.startConsumer()
	}
}

func (logWriter *RedisLogWriter) startConsumer() {
	for logLine := range logWriter.inputChannel {
		err := logWriter.redisClient.RPush(logWriter.logName, logLine).Err()
		if err != nil {
			panic(err)
		}
	}
}
