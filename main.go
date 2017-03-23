package main

import (
	"hello/models"
	_ "hello/routers"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	msg, valid := models.ParseTextMsg("&&JPY")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true
	beego.Run()
}
