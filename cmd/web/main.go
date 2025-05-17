package main

import (
	"net/http"

	"github.com/jgr142/portfolio/internal/core"
	"github.com/jgr142/portfolio/internal/web"
)

func main() {
	core := core.New(nil)

	handlers := web.InitHandlers(core.Logger)
	mux := web.InitMux(handlers)

	core.Logger.Info("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	core.Logger.Error(err.Error())
}
