package config

type Config struct {
	PgConn string
}

func New() Config {
	return Config{
		PgConn: "user=postgres password=7070 dbname=tn-bot port=5432 sslmode=disable",
	}
}
