package main

import (
	"hello/models"
	_ "hello/routers"
	"log"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	msg, _ := models.SpliteTextMsg("&&JPY")
	log.Print(msg)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true
	beego.Run()
}
