package handler

import "net/http"

func NewMux(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", h.Ping)
	mux.HandleFunc("/aso-msg", h.SendMessage)
	middlewares := []MiddleWare{
		CORS,
		Recovery,
		LoggerWithFormatter,
	}
	return Use(mux, middlewares...)
}
