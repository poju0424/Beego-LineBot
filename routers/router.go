package routers

import (
	"hello/controllers"

	"hello/models"

	"github.com/astaxie/beego"
)

// var bot *linebot.Client

// type myHandler struct{}

// func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func init() {
	beego.Router("/", &controllers.MainController{})

	// var err error
	// bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	// log.Println("Bot:", bot, " err:", err)
	beego.Handler("/callback", &models.LineHandler{})
}
