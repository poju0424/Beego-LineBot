package controllers

import (
	"fmt"
	"hello/Util/Debug"

	"github.com/astaxie/beego"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineBotController struct {
	beego.Controller
}

type LineMsg struct {
	Id              string `form:"-"`
	ContentType     int    `form:"username"`
	From            string `form:"age"`
	CreatedTime     int
	To              []string
	ToType          int
	ContentMetadata interface{}
	text            string
	location        interface{}
}

var bot *linebot.Client

func (c *LineBotController) Post() {
	msg := LineMsg{}
	if err := c.ParseForm(&msg); err != nil {
		Debug.CheckErr(err)
	}
	fmt.Print(msg)
	//func callbackHandler(w http.ResponseWriter, r *http.Request) {
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
