package config

import "fmt"

type ClientsConfig struct {
	WriterDbURL string
	ReaderDbURL string
}

type AppConfig struct {
	Address string
}

type Config struct {
	App     AppConfig
	Clients ClientsConfig
}

func New() *Config {
	c := new(Config)

	c.App = AppConfig{
		Address: fmt.Sprintf(":%d", LookupEnv("PORT", 7000)),
	}

	c.Clients = ClientsConfig{
		WriterDbURL: LookupEnv("DB_WRITER_URL", ""),
		ReaderDbURL: LookupEnv("DB_READER_URL", ""),
	}

	return c
}
