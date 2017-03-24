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

	code, name := models.FuzzySearch("123")
	log.Print(code + name)
	log.Print("code + name")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true
	beego.Run()
}
