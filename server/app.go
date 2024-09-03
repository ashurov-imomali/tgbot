package server

import (
	"github.com/ashurov-imomali/tgbot/config"
	"github.com/ashurov-imomali/tgbot/handler"
	"github.com/ashurov-imomali/tgbot/logger"
	"github.com/ashurov-imomali/tgbot/usecase"
	"net/http"
)

type App struct {
	h *http.Server
}

func New() *App {
	l := logger.New()
	conf, err := config.GetConfigs()
	if err != nil {
		l.Fatal(err)
	}
	useCase := usecase.New(conf, l)
	h := handler.New(useCase, l)
	mux := handler.NewMux(h)
	l.Printf("startet listen addres %s", conf.Address)
	return &App{
		h: &http.Server{
			Addr:    conf.Address,
			Handler: mux,
		},
	}

}

func (a App) Run() error {
	return a.h.ListenAndServe()
}
