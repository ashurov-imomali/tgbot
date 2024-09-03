package handler

import (
	"github.com/ashurov-imomali/tgbot/logger"
	"github.com/ashurov-imomali/tgbot/usecase"
	"net/http"
	"time"
)

type Handler struct {
	u usecase.IUseCase
	l logger.ILogger
}

func New(u usecase.IUseCase, l logger.ILogger) *Handler {
	return &Handler{l: l, u: u}
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	//pong := h.u.Pong()
	time.Sleep(3 * time.Second)
	w.WriteHeader(500)
}
