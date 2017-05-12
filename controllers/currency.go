package controllers

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
)

type CurrencyController struct {
	beego.Controller
}

func (c *CurrencyController) Get() {
	name := c.Ctx.Input.Param(":name")
	time := c.Ctx.Input.Param(":time")
	// url := "http://rate.bot.com.tw/xrt/quote/" + time + "/" + name + ""
	// log.Print(url)

	data := getData(time, name)
	c.Data["Body"] = data
	c.TplName = "index.html"
}

func getData(time, name string) interface{} {
	url := "http://rate.bot.com.tw/xrt/quote/" + time + "/" + name + ""
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Print(err)
	}
	table := doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		log.Print(i)
		// log.Print(s.ClosestSelection(s.Find("td")))
		// log.Print(s.Length())
		log.Print(s.Find("td").Length())
		log.Print(s.Find("td").Eq(0).Text())

	})

	log.Print(55665566)
	// log.Print(table.Text())

	return table.Text()
}
