package controllers

import (
	"encoding/json"
	"fmt"

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
	// msg := LineMsg{}
	// if err := c.ParseForm(&msg); err != nil {
	// 	Debug.CheckErr(err)
	// }
	// fmt.Print(msg)
	var ob LineMsg
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	// objectid := models.AddOne(ob)
	// this.Data["json"] = map[string]interface{}{"ObjectId": objectid}
	// this.ServeJSON()
	fmt.Print(ob)

}
