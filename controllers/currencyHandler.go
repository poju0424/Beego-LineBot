package controllers

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	chart "github.com/wcharczuk/go-chart"
)

type CurrencyHandler struct{}

type RateHistoryStruct struct {
	CashBuy      []float64
	CashSell     []float64
	RateBuy      []float64
	RateSell     []float64
	Date         []string
	CurrencyName string
}

func (*CurrencyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	params := strings.Split(r.RequestURI, "/")
	if len(params) == 4 {
		time := params[2]
		name := params[3]
		data := getData(time, name)
		createChart(data)
		w.Write([]byte(data.CurrencyName))
		return
	}

}

func NewRateHistoryStruct(name string) *RateHistoryStruct {
	obj := new(RateHistoryStruct)
	obj.CurrencyName = name
	return obj
}

func getData(time, name string) *RateHistoryStruct {
	url := "http://rate.bot.com.tw/xrt/quote/" + time + "/" + name + ""
	doc, err := goquery.NewDocument(url)
	history := NewRateHistoryStruct(name)

	if err != nil {
		log.Print(err)
	}
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		date := s.Find("td").Eq(0).Text()
		cashBuy, _ := strconv.ParseFloat(s.Find("td").Eq(2).Text(), 64)
		cashSell, _ := strconv.ParseFloat(s.Find("td").Eq(3).Text(), 64)
		rateBuy, _ := strconv.ParseFloat(s.Find("td").Eq(4).Text(), 64)
		rateSell, _ := strconv.ParseFloat(s.Find("td").Eq(5).Text(), 64)

		history.CashBuy = append(history.CashBuy, cashBuy)
		history.CashSell = append(history.CashSell, cashSell)
		history.RateBuy = append(history.RateBuy, rateBuy)
		history.RateSell = append(history.RateSell, rateSell)
		history.Date = append(history.Date, date)
	})
	return history
}

func createChart(data *RateHistoryStruct) {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: data.CashBuy,
				YValues: data.CashSell,
			},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	log.Print(err)
}
