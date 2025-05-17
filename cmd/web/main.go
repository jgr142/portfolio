package main

import (
	"net/http"
	"os"

	"github.com/jgr142/portfolio/internal/infra/db"
	"github.com/jgr142/portfolio/internal/platform"
	"github.com/jgr142/portfolio/internal/web"
)

func main() {
	p := platform.New(nil)

	sql, err := db.Open("web:MaddenPro_101@/portfolio?parseTime=true")
	if err != nil {
		p.Logger.Error(err.Error())
		os.Exit(1)
	}

	defer sql.Close()

	dal := db.New(sql)

	handlers := web.InitHandlers(dal, p.Logger)
	mux := web.InitMux(handlers)

	p.Logger.Info("starting server on :4000")

	err = http.ListenAndServe(":4000", mux)
	p.Logger.Error(err.Error())
}
