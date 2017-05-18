package controllers

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"

	"time"

	"github.com/PuerkitoBio/goquery"
	chart "github.com/wcharczuk/go-chart"
)

type CurrencyHandler struct{}

type RateHistoryStruct struct {
	CashBuy      []float64
	CashSell     []float64
	RateBuy      []float64
	RateSell     []float64
	Date         []time.Time
	CurrencyName string
}

func (*CurrencyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	params := strings.Split(r.RequestURI, "/")
	if len(params) == 4 {
		time := params[2]
		name := params[3]
		data := getData(time, name)
		buff := createChart(data)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
		if _, err := w.Write(buff.Bytes()); err != nil {
			log.Print(err)
		}
		return
	}

}

func NewRateHistoryStruct(name string) *RateHistoryStruct {
	obj := new(RateHistoryStruct)
	obj.CurrencyName = name
	return obj
}

func getData(date, name string) *RateHistoryStruct {
	url := "http://rate.bot.com.tw/xrt/quote/" + date + "/" + name + ""
	doc, err := goquery.NewDocument(url)
	history := NewRateHistoryStruct(name)

	if err != nil {
		log.Print(err)
	}
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		date := s.Find("td").Eq(0).Text()
		date1, _ := time.Parse("2006/01/02", date)

		cashBuy, _ := strconv.ParseFloat(s.Find("td").Eq(2).Text(), 64)
		cashSell, _ := strconv.ParseFloat(s.Find("td").Eq(3).Text(), 64)
		rateBuy, _ := strconv.ParseFloat(s.Find("td").Eq(4).Text(), 64)
		rateSell, _ := strconv.ParseFloat(s.Find("td").Eq(5).Text(), 64)

		history.CashBuy = append(history.CashBuy, cashBuy)
		history.CashSell = append(history.CashSell, cashSell)
		history.RateBuy = append(history.RateBuy, rateBuy)
		history.RateSell = append(history.RateSell, rateSell)
		history.Date = append(history.Date, date1)
	})
	return history
}

func createChart(data *RateHistoryStruct) *bytes.Buffer {
	graph := chart.Chart{
		Width:  1024,
		Height: 512,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeValueFormatterWithFormat("2006/01/02"),
			Ticks:          setTicks(data.Date),
			TickStyle:      chart.Style{Show: true},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: data.Date,
				YValues: data.CashSell,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Print(err)
	}
	return buffer
}

// func (mhr *MarketHoursRange) makeTicks(vf ValueFormatter, times []time.Time) []Tick {
// 	ticks := make([]Tick, len(times))
// 	for index, t := range times {
// 		ticks[index] = Tick{
// 			Value: util.Time.ToFloat64(t),
// 			Label: vf(t),
// 		}
// 	}
// 	return ticks
// }

func setTicks(times []time.Time) []chart.Tick {
	ticks := make([]chart.Tick, len(times))
	for index, t := range times {
		ticks[index] = chart.Tick{
			Value: chart.Time.ToFloat64(t),
			Label: t.String(),
		}
	}
	log.Print(ticks)
	return ticks
}
