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
	mux.HandleFunc("GET /{$}", r.h.home)
	mux.HandleFunc("GET /project/view/{id}", r.h.projectView)
	mux.HandleFunc("GET /project/create", r.h.projectCreate)
	mux.HandleFunc("POST /snippet/create", r.h.snippetCreatePost)

	return mux
}
