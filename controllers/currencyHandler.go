package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type CurrencyHandler struct{}

type PerHistory struct {
	date     string
	cashBuy  float64
	cashSell float64
	rateBuy  float64
	rateSell float64
}

type RateHistoryStruct struct {
	Items        []PerHistory
	CurrencyName string
}

func (*CurrencyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RequestURI)

}

func NewRateHistoryStruct(name string) *RateHistoryStruct {
	obj := new(RateHistoryStruct)
	obj.CurrencyName = name
	return obj
}

func (box *RateHistoryStruct) AddItem(item PerHistory) []PerHistory {
	box.Items = append(box.Items, item)
	return box.Items
}

func getData(time, name string) interface{} {
	url := "http://rate.bot.com.tw/xrt/quote/" + time + "/" + name + ""
	doc, err := goquery.NewDocument(url)
	history := NewRateHistoryStruct(name)
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
			cashBuy:  cashBuy,
			cashSell: cashSell,
			rateBuy:  rateBuy,
			rateSell: rateSell,
		}
		history.AddItem(perHistory)
	})
	return history
}
