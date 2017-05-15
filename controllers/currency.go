package controllers

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
)

type CurrencyController struct {
	beego.Controller
}

type PerHistory struct {
	date     string
	name     string
	cashBuy  float64
	cashSell float64
	rateBuy  float64
	rateSell float64
}

type RateHistoryStruct struct {
	Items []PerHistory
}

func (c *CurrencyController) Get() {
	name := c.Ctx.Input.Param(":name")
	time := c.Ctx.Input.Param(":time")

	data := getData(time, name)
	c.Data["Body"] = data
	c.TplName = "index.html"
}

func NewRateHistoryStruct() *RateHistoryStruct {
	obj := new(RateHistoryStruct)
	return obj
}

func (box *RateHistoryStruct) AddItem(item PerHistory) []PerHistory {
	box.Items = append(box.Items, item)
	return box.Items
}

func getData(time, name string) interface{} {
	url := "http://rate.bot.com.tw/xrt/quote/" + time + "/" + name + ""
	doc, err := goquery.NewDocument(url)
	history := NewRateHistoryStruct()
	if err != nil {
		log.Print(err)
	}
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		cashBuy, _ := strconv.ParseFloat(s.Find("td").Eq(2).Text(), 64)
		cashSell, _ := strconv.ParseFloat(s.Find("td").Eq(3).Text(), 64)
		rateBuy, _ := strconv.ParseFloat(s.Find("td").Eq(4).Text(), 64)
		rateSell, _ := strconv.ParseFloat(s.Find("td").Eq(5).Text(), 64)

		perHistory := PerHistory{
			date:     s.Find("td").Eq(0).Text(),
			name:     s.Find("td").Eq(1).Text(),
			cashBuy:  cashBuy,
			cashSell: cashSell,
			rateBuy:  rateBuy,
			rateSell: rateSell,
		}
		history.AddItem(perHistory)
	})
	return history
}
