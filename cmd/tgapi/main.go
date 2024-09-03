package main

import (
	"github.com/ashurov-imomali/tgbot/server"
	"log"
)

func main() {
	app := server.New()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
