package handlers

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Handle(w http.ResponseWriter, r *http.Request, bot *linebot.Client) {
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
}
