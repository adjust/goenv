package goenv

import (
	"testing"
)

func TestGetAmqp(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "amqp", "nil")
	if goenv.GetAmqp() != "amqp://uhe:roh@hih:870/" {
		t.Error("amqp != amqp://uhe:roh@hih:870/")
	}
}

func TestGetAmqpNotFound(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "nonexistent", "nil")
	if goenv.GetAmqp() != "amqp://guest:guest@localhost:5672/" {
		t.Error("amqp != amqp://guest:guest@localhost:5672/")
	}
}

func TestGetNamedAmqp(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "amqp", "nil")
	if goenv.GetNamedAmqp("custom") != "amqp://rtn:kbo@aar:473/" {
		t.Error("namedAmqp != amqp://rtn:kbo@aar:473/")
	}
}

func TestGetNamedAmqpNotFound(t *testing.T) {
	goenv := NewGoenv("./test_config.yml", "amqp", "nil")
	if goenv.GetNamedAmqp("nonexistent") != "amqp://guest:guest@localhost:5672/" {
		t.Error("nonexistent != amqp://guest:guest@localhost:5672/")
	}
}
