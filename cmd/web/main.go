package main

import (
	"net/http"

	"github.com/jgr142/portfolio/internal/platform"
	"github.com/jgr142/portfolio/internal/web"
)

func main() {
	platform := platform.New(nil)

	handlers := web.InitHandlers(platform.Logger)
	mux := web.InitMux(handlers)

	platform.Logger.Info("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	platform.Logger.Error(err.Error())
}
