package goenv

func (goenv *Goenv) GetPort() string {
	port := goenv.Get("port", "8080")
	return port
}
