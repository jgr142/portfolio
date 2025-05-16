package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := routes()

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("starting server on :4000")

	err := http.ListenAndServe(":4000", routes())
	logger.Error(err.Error())
}
