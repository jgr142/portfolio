package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/jgr142/portfolio/internal/infra/db"
	"github.com/jgr142/portfolio/internal/platform"
	"github.com/jgr142/portfolio/internal/web"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:MaddenPro_101@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	p := platform.New(nil)

	sql, err := db.Open(*dsn)
	if err != nil {
		p.Logger.Error(err.Error())
		os.Exit(1)
	}

	defer sql.Close()

	dal := db.New(sql)

	tCache, err := web.NewTemplateCache()
	if err != nil {
		p.Logger.Error(err.Error())
		os.Exit(1)
	}

	handlers := web.InitHandlers(dal, tCache, p.Logger)
	mux := web.InitMux(handlers)

	p.Logger.Info("starting server on %s", slog.String("addr", *addr))

	err = http.ListenAndServe(*addr, mux)
	p.Logger.Error(err.Error())
}
