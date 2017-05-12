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
	historys []PerHistory
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
	var history *RateHistoryStruct
	if err != nil {
		log.Print(err)
	}
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		// log.Print(i)

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
		// log.Print(s.ClosestSelection(s.Find("td")))
		// log.Print(s.Length())
		// log.Print(s.Find("td").Length())
		// log.Print(s.Find("td").Eq(0).Text()) //date
		// log.Print(s.Find("td").Eq(1).Text()) //currency
		// log.Print(s.Find("td").Eq(2).Text()) //cashbuy
		// log.Print(s.Find("td").Eq(3).Text()) //cashsell
		// log.Print(s.Find("td").Eq(4).Text()) //ratebuy
		// log.Print(s.Find("td").Eq(5).Text()) //ratesell

	})

	log.Print(55665566)
	log.Print(history)
	// log.Print(table.Text())

	return history
}

func (box *RateHistoryStruct) AddItem(item PerHistory) []PerHistory {
	box.historys = append(box.historys, item)
	return box.historys
}
