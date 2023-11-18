package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/protomem/socnet/pkg/closing"
	"github.com/protomem/socnet/user-service/internal/config"
)

type App struct {
	conf config.Config

	server *http.Server

	closer *closing.Closer
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
		closer: closing.New(),
	}
}

func (app *App) Run() error {
	ctx := context.Background()

	app.registerOnShutdown()

	errs := make(chan error)

	go func() { errs <- app.startServer() }()
	go func() { errs <- app.gracefullShutdown(ctx) }()

	err := <-errs
	if err != nil {
		return err
	}

	return nil
}

func (app *App) registerOnShutdown() {
	app.closer.Add(app.server.Shutdown)
}

func (app *App) startServer() error {
	err := app.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (app *App) gracefullShutdown(ctx context.Context) error {
	<-wait()

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := app.closer.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

func wait() <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	return ch
}
