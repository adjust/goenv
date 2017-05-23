package goenv

import (
	"github.com/adjust/redis"

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
	}

	logWriter.redisClient = redis.NewClient(options)
	err = logWriter.redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}

	logWriter.inputChannel = make(chan string, 10000)
	go logWriter.startConsumer()

	return logWriter, nil
}

func (logWriter *RedisLogWriter) Write(p []byte) (n int, err error) {
	logWriter.inputChannel <- string(p)
	return len(p), nil
}

func (writer *RedisLogWriter) startConsumer() {
	var todo int
	var batch []string

	for logLine := range writer.inputChannel {
		batch = append(batch, logLine)

		// are we still on a batch run?
		if todo > 0 {
			todo--
			continue
		}

		// batch run done, flushing
		writer.pushToRedis(batch)
		batch = []string{}

		// fetch next batch run size
		todo = len(writer.inputChannel) - 1 // otherwise we'd lose the last line
	}
}

func (writer *RedisLogWriter) pushToRedis(logLines []string) {
	err := writer.redisClient.RPush(writer.logName, logLines...).Err()
	if err != nil {
		panic(err)
	}
}
