package models

import (
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

type LineHandler struct{}

func (*LineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	// log.Println("Bot:", bot, " err:", err)

	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg, isValid := spliteTextMsg(message.Text)
				// log.Print(msg)
				if isValid {
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(GetRateInfo(msg))).Do(); err != nil {
					// 	log.Print(err)
					// }
					if _, err = bot.ReplyMessage(event.ReplyToken, ReplyTemplateMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}

func spliteTextMsg(msg string) (subMsg string, isValid bool) {
	re := regexp.MustCompile("^&&(.*)")
	arr := re.FindStringSubmatch(msg)
	if len(arr) > 0 {
		isValid = true
		subMsg = arr[1]
	} else {
		isValid = false
	}
	return
}
