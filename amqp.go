package goenv

import (
	"log"
)

func GetAmqp() string {
	result := GetNamedAmqp("amqp")
	return result
}

func GetNamedAmqp(name string) string {
	user := Get(name+".user", "guest")
	pass := Get(name+".pass", "guest")
	host := Get(name+".host", "localhost")
	port := Get(name+".port", "5672")
	result := "amqp://" + user + ":" + pass + "@" + host + ":" + port + "/"
	return result
}
