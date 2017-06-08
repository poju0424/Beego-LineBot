package controllers

import (
	"encoding/json"
	"hello/service"
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
				log.Print(message.Text)
				msg := message.Text
				if _, err = bot.ReplyMessage(event.ReplyToken, replyTemplateMessage(msg)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, service.GetNerybyBank(message.Latitude, message.Longitude)).Do(); err != nil {
					log.Print(err)
				}
			}
		} else if event.Type == linebot.EventTypePostback {
			byteArr := []byte(event.Postback.Data)
			var postbackObj service.PostbackObj
			err := json.Unmarshal(byteArr, &postbackObj)
			if err != nil {
				log.Print(err)
			}
			switch postbackObj.Method {
			case "text":
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(postbackObj.Data)).Do()
			case "image":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(postbackObj.Data, postbackObj.Data)).Do(); err != nil {
					log.Print(err)
				}
			}

		}
	}
}

func spliteTextMsg(msg string) (subMsg string, isValid bool) { //only response text start with "&&"
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

func replyTemplateMessage(request string) (templateMsg linebot.Message) {
	content, name := service.GetRateInfo(request)
	code, _ := service.FuzzySearch(request)
	var AltText = content
	if len(content) <= 0 || len(name) <= 0 {
		return nil
	}
	template := linebot.NewButtonsTemplate(
		"", "", content,
		linebot.NewURITemplateAction("台銀網站", "https://goo.gl/ZCXw47"),
		linebot.NewPostbackTemplateAction("近3個月現金匯率", service.NewJString("image", "https://beegolinebot.herokuapp.com/currency/ltm/"+code+""), ""),
		linebot.NewPostbackTemplateAction("附近的分行", service.NewJString("text", "請傳送位置資訊給我"), ""),
		linebot.NewMessageTemplateAction("重新查詢", name),
	)

	templateMsg = linebot.NewTemplateMessage(AltText, template)
	return
}
