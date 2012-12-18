package goenv

func GetAmqp() string {
	return GetNamedAmqp("amqp")
}

func GetNamedAmqp(name string) string {
	user := Get("amqp.user", "guest")
	pass := Get("amqp.pass", "guest")
	host := Get("amqp.host", "localhost")
	port := Get("amqp.port", "5672")
	result := "amqp://" + user + ":" + pass + "@" + host + ":" + port + "/"
	return result
}
