package routers

import (
	"hello/controllers"
	"log"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/callback", &controllers.LineBotController{})
	// s.RegisterCodec(json.NewCodec(), "application/json")
	// s.RegisterService(new(HelloService), "")
	// beego.Handler("/rpc", s)
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	// var handle http.Handler

	beego.Handler("/callback", &myHandler{})
}

// func callbackHandler(w http.ResponseWriter, r *http.Request) {
// 	events, err := bot.ParseRequest(r)

// 	if err != nil {
// 		if err == linebot.ErrInvalidSignature {
// 			w.WriteHeader(400)
// 		} else {
// 			w.WriteHeader(500)
// 		}
// 		return
// 	}

// 	for _, event := range events {
// 		if event.Type == linebot.EventTypeMessage {
// 			switch message := event.Message.(type) {
// 			case *linebot.TextMessage:
// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
// 					log.Print(err)
// 				}
// 			}
// 		}
// 	}
// }
