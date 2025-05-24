package web

import (
	"log/slog"
	"net/http"
)

func (h *Handler) serverError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	h.logger.Error(
		err.Error(),
		slog.String("method", method),
		slog.String("uri", uri),
	)

	http.Error(w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

func (h *Handler) clientError(
	w http.ResponseWriter,
	status int,
) {
	http.Error(w, http.StatusText(status), status)
}
