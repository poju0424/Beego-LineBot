package controllers

import (
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

func (this *LineBotController) Post() {
	// var err error
	// bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	// log.Println("Bot:", bot, " err:", err)
	// events, err := bot.ParseRequest(r)
	// fmt.Print(events)
	// fmt.Print("5566")
	// fmt.Print(this.Ctx.Input.Param("result"))
	// fmt.Print("5566")

}
