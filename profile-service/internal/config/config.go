package config

import "os"

type Config struct {
	HTTP struct {
		Addr string
	}
}

func New() (Config, error) {
	var (
		conf Config
		ok   bool
	)

	conf.HTTP.Addr, ok = os.LookupEnv("HTTP__ADDR")
	if !ok {
		conf.HTTP.Addr = ":8080"
	}

	return conf, nil
}
