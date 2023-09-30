package app

import (
	"fmt"
	"net/http"

	"github.com/protomem/socnet/profile-service/internal/config"
)

type App struct {
	conf config.Config

	server *http.Server
}

func New(conf config.Config) *App {
	return &App{
		conf: conf,
		server: &http.Server{
			Addr: conf.HTTP.Addr,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Profile Service v0.1.0")
			}),
		},
	}
}

func (app *App) Run() error {
	return app.server.ListenAndServe()
}
