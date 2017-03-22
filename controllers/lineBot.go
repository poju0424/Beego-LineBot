package controllers

import (
	"github.com/astaxie/beego"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineBotController struct {
	beego.Controller
}

var bot *linebot.Client

func (c *LineBotController) Get() {
	//git@github.com:poju0424/Beego-LineBot.git
	// events, err := bot.ParseRequest(r)
	// if err != nil {
	// 	if err == linebot.ErrInvalidSignature {
	// 		w.WriteHeader(400)
	// 	} else {
	// 		w.WriteHeader(500)
	// 	}
	// 	return
	// }
	// for _, event := range events {
	// 	switch message := event.Message.(type) {
	// 	case *linebot.TextMessage:
	// 		bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
	// 	}
	// }
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}
