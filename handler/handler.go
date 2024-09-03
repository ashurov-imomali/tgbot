package handler

import (
	"encoding/json"
	"github.com/ashurov-imomali/tgbot/logger"
	"github.com/ashurov-imomali/tgbot/models"
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

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Msg
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		h.l.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := h.u.SendMessageToGroup(msg.Msg); err != nil {
		h.l.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

}
