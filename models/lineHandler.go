package models

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

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
				if isValid {
					if _, err = bot.ReplyMessage(event.ReplyToken, ReplyTemplateMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				}
			case *linebot.LocationMessage:
				getNerybyBank(message.Latitude, message.Longitude)
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

func getNerybyBank(lat, lon float64) {
	log.Print(lat)
	log.Print(lon)
	latitude := strconv.FormatFloat(lat, 'f', -1, 64)
	longitude := strconv.FormatFloat(lon, 'f', -1, 64)
	name := "臺灣銀行股份有限公司"
	APIKey := os.Getenv("GoogleMapNearbySearchKey")
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + latitude + "," + longitude + "&name=" + name + "&key=" + APIKey + "&language=zh-TW&types=bank&rankby=distance"
	log.Print(url)
}
