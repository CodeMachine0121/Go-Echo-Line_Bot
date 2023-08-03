package main

import (
	handlers "go-line/Handlers"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {

	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	channel_secret := os.Getenv("CHANNEL_SECRET")

	bot, err := linebot.New(channel_secret, channel_access_token)

	if err != nil {
		log.Fatal("error occurred while creating")
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		handlers.Handle(w, r, bot)
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
