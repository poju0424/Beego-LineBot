package controllers

import (
	"Beego-LineBot/models"
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	chart "github.com/wcharczuk/go-chart"
)

type CurrencyHandler struct{}

func (*CurrencyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.RequestURI, "/")
	if len(params) == 4 {
		time := params[2]
		name := params[3]
		data := models.getTicksIntervalArgs.GetData(time, name)
		models.
		buff := createChart(data)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
		if _, err := w.Write(buff.Bytes()); err != nil {
			log.Print(err)
		}
		return
	}

}

func createChart(data *service.RateHistoryStruct) *bytes.Buffer {
	graph := chart.Chart{
		Title: data.CurrencyName + "(" + data.Date[0].Format("Jan 2 2006") + ")",
		TitleStyle: chart.Style{
			Show: true,
		},
		Width:  1024,
		Height: 512,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeValueFormatterWithFormat("2006/01/02"),
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Ticks: makeTicks(data),
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 60,
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Name:    string("CashSell"),
				XValues: data.Date,
				YValues: data.CashSell,
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetAlternateColor(4),
					StrokeWidth: 5,
				},
			},
			chart.TimeSeries{
				Name:    string("CashBuy"),
				XValues: data.Date,
				YValues: data.CashBuy,
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetAlternateColor(6),
					StrokeWidth: 5,
				},
			},
		},
	}

	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Print(err)
	}
	return buffer
}

func makeTicks(data *service.RateHistoryStruct) (ticks []chart.Tick) {
	min, max := findSliceMinMax(data)
	scale, interval, fixed := getTicksIntervalArgs(max)
	dMax := decimal.NewFromFloat(max).Mul(decimal.NewFromFloat(scale)).Ceil().Div(decimal.NewFromFloat(scale))
	dMin := decimal.NewFromFloat(min).Mul(decimal.NewFromFloat(scale)).Floor().Div(decimal.NewFromFloat(scale))

	for dMax.GreaterThanOrEqual(dMin) {
		f64Val, _ := dMax.Float64()
		temp := chart.Tick{Value: f64Val, Label: dMax.StringFixed(fixed)}
		ticks = append(ticks, temp)
		dMax = dMax.Add(decimal.NewFromFloat(interval))
	}
	return
}

func getTicksIntervalArgs(input float64) (scale, interval float64, fixed int32) {
	length := len(strconv.FormatFloat(input, 'f', 0, 64))
	if length <= 1 {
		interval = (-0.005)
		scale = 200
		fixed = 3
	} else {
		interval = (-0.1)
		scale = 10
		fixed = 1
	}
	return
}

func findSliceMinMax(v *service.RateHistoryStruct) (min, max float64) {
	if len(v.CashSell) > 0 && len(v.CashBuy) > 0 {
		min = v.CashBuy[0]
		max = v.CashSell[0]
	}
	for i := 1; i < len(v.CashSell); i++ {
		if v.CashBuy[i] < min {
			min = v.CashBuy[i]
		}
		if v.CashSell[i] > max {
			max = v.CashSell[i]
		}
	}
	return
}
