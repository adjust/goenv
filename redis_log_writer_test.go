package goenv

import (
	"testing"
)

func TestRedisLogWriter(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "redis", "nil")
	writer, err := goenv.NewRedisLogWriter("test123")
	if err != nil {
		t.Fatalf("couldn't connect to redis for log writer test: %s", err)
	}
	writer.Write([]byte("dingdong"))
	value, err := writer.redisClient.LPop("test123").Result()
	if err != nil {
		t.Errorf("error reading log writer: %s", err)
	}

	if value != "dingdong" {
		t.Error("output didn't match input from log writer")
	}
}
