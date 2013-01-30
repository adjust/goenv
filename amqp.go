package goenv

import (
	"fmt"
)

func (goenv *Goenv) GetAmqp() string {
	result := goenv.GetNamedAmqp("amqp")
	return result
}

func (goenv *Goenv) GetNamedAmqp(name string) string {
	user := goenv.Get(name+".user", "guest")
	pass := goenv.Get(name+".pass", "guest")
	host := goenv.Get(name+".host", "localhost")
	port := goenv.Get(name+".port", "5672")
	result := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	return result
}
