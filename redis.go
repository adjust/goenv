package goenv

func GetRedis() (host, port string, db int) {
	return GetNamedRedis("redis")
}

func GetNamedRedis(name string) (host, port string, db int) {
	host = Get(name+".host", "localhost")
	port = Get(name+".port", "6379")
	db = GetInt(name+".db", 0)

	return
}
