package goenv

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestRedisLogWriter(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "config", "nil")
	writer, err := goenv.NewRedisLogWriter("test123")
	if err != nil {
		t.Fatalf("couldn't connect to redis for log writer test: %s", err)
	}
	writer.Write([]byte("dingdong"))
	time.Sleep(10 * time.Millisecond)
	value, err := writer.redisClient.LPop("test123").Result()
	if err != nil {
		t.Errorf("error reading log writer: %s", err)
	}

	if value != "dingdong" {
		t.Error("output didn't match input from log writer")
	}
}

func TestRedisLogWriterForStdLog(t *testing.T) {
	goenv := NewGoenv("./test_config2.yml", "config", "")
	log.Println("hello world")

	writer, err := goenv.NewRedisLogWriter("")

	if err != nil {
		t.Fatalf("couldn't connect to redis for log writer test: %s", err)
	}
	value, err := writer.redisClient.LPop("test123.log").Result()
	if err != nil {
		t.Errorf("error reading log writer: %s", err)
	}

	if !strings.Contains(value, "hello world") {
		t.Errorf("output didn't match input from log writer: %s", value)
	}
}
