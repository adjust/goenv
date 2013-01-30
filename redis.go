package goenv

func (goenv *Goenv) GetRedis() (host, port string, db int) {
	return goenv.GetNamedRedis("redis")
}

func (goenv *Goenv) GetNamedRedis(name string) (host, port string, db int) {
	host = goenv.Get(name+".host", "localhost")
	port = goenv.Get(name+".port", "6379")
	db = goenv.GetInt(name+".db", 0)

	return
}
