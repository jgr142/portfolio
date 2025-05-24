package web

import "net/http"

type Router struct {
	h Handler
}

func InitMux(h Handler) http.Handler {
	router := Router{h: h}
	return router.routes()
}

func (r *Router) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", r.h.home)
	mux.HandleFunc("GET /project/view/{id}", r.h.projectView)
	mux.HandleFunc("GET /project/create", r.h.projectCreate)
	mux.HandleFunc("POST /snippet/create", r.h.projectCreatePost)

	return mux
}
