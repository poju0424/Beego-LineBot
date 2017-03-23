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
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(GetRateInfo(message.Text))).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func ParseTextMsg(msg string) (subMsg string, isValid bool) {
	var re = regexp.MustCompile(`^\&&(.*)`).Split(msg, 5)
	log.Print(re)
	// Debug.CheckErr(err)
	// log.Print(matched)
}
