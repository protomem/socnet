package main

import (
	"github.com/protomem/socnet/user-service/internal/app"
	"github.com/protomem/socnet/user-service/internal/config"
)

func main() {
	var err error

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	err = app.New(conf).Run()
	if err != nil {
		panic(err)
	}
}
