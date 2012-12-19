package goenv

import (
	"github.com/adeven/goenv"
	"testing"
)

func TestGetAmqp(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("amqp")
	if goenv.GetAmqp() != "amqp://uhe:roh@hih:870/" {
		t.Error("amqp != amqp://uhe:roh@hih:870/")
	}
}

func TestGetAmqpNotFound(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("nonexistent")
	if goenv.GetAmqp() != "amqp://guest:guest@localhost:5672/" {
		t.Error("amqp != amqp://guest:guest@localhost:5672/")
	}
}

func TestGetNamedAmqp(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("amqp")
	if goenv.GetNamedAmqp("custom") != "amqp://rtn:kbo@aar:473/" {
		t.Error("namedAmqp != amqp://rtn:kbo@aar:473/")
	}
}

func TestGetNamedAmqpNotFound(t *testing.T) {
	goenv.SetConfigFile("./config/config.yaml")
	goenv.SetEnvironment("amqp")
	if goenv.GetNamedAmqp("nonexistent") != "amqp://guest:guest@localhost:5672/" {
		t.Error("nonexistent != amqp://guest:guest@localhost:5672/")
	}
}
