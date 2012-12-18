package goenv

func GetPostgres() string {
	return GetNamedPostgres("postgres")
}

func GetNamedPostgres(name string) string {
	user := Get(name+".user", "postgres")
	host := Get(name+".host", "localhost")
	dbst := Get(name+".db", "0")
	result := "user=" + user + " dbname=" + dbst + " sslmode=disable host=" + host
	return result
}
