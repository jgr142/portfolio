package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := routes()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	logger.Error(err.Error())
}
