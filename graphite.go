package goenv

func (goenv *Goenv) GetGraphite() (host string, port int) {
	return goenv.GetNamedGraphite("graphite")
}

func (goenv *Goenv) GetNamedGraphite(name string) (host string, port int) {
	host = goenv.Get(name+".host", "")
	port = goenv.GetInt(name+".port", 2003)

	return
}
