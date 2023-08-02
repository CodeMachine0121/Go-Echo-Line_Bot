package main

import (
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
		events, err := bot.ParseRequest(r)
		if err != nil {
			log.Fatal(err)
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch msg := event.Message.(type) {

				case *linebot.TextMessage:
					_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg.Text)).Do()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
