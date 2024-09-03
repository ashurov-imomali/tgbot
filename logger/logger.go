package logger

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	l zerolog.Logger
}

func New() Logger {
	return Logger{
		zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(3).Logger(),
	}
}

func (k Logger) Println(a ...interface{}) {
	for _, v := range a {
		if indent, err := json.MarshalIndent(v, "", " "); err == nil {
			k.l.Print(string(indent))
		}
	}
}

func (k Logger) Printf(format string, a ...interface{}) {
	k.l.Printf(format, a...)
}

func (k Logger) Error(a ...interface{}) {
	k.l.Error().Msgf("%v", a)
}

func (k Logger) Fatal(a ...interface{}) {
	k.l.Fatal().Msgf("%v", a)
}

type ILogger interface {
	Println(a ...interface{})
	Printf(format string, a ...interface{})
	Error(a ...interface{})
	Fatal(a ...interface{})
}
