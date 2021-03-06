package model

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func GetNerybyBank(lat, lon float64) (templateMsg linebot.Message) {
	latitude := strconv.FormatFloat(lat, 'f', -1, 64)
	longitude := strconv.FormatFloat(lon, 'f', -1, 64)
	name := "臺灣銀行股份有限公司"
	APIKey := os.Getenv("GoogleMapNearbySearchKey")
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + latitude + "," + longitude + "&name=" + name + "&key=" + APIKey + "&language=zh-TW&type=bank&rankby=distance"

	type LatLng struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	type Bounds struct {
		Northeast LatLng `json:"northeast"`
		Southwest LatLng `json:"southwest"`
	}

	type Geometry struct {
		Location LatLng `json:"location"`
		Viewport Bounds `json:"viewport"`
	}

	type Photos struct {
		Photo_reference string `json:"photo_reference"`
	}

	type Result struct {
		Name     string   `json:"name"`
		Vicinity string   `json:"vicinity"`
		Photos   []Photos `json:"photos"`
		Geometry Geometry `json:"geometry"`
		Place_id string   `json:"place_id"`
	}

	type Results struct {
		Results []Result `json:"results"`
		Status  string   `json:"status"`
	}

	nearby := new(Results)
	err := getJSON(url, nearby)
	if err != nil {
		panic(err)
	}

	var s []*linebot.CarouselColumn
	if nearby.Status == "OK" {
		for i := 0; i < 5; i++ {
			f1 := strconv.FormatFloat(nearby.Results[i].Geometry.Location.Lat, 'f', -1, 64)
			f2 := strconv.FormatFloat(nearby.Results[i].Geometry.Location.Lng, 'f', -1, 64)
			destination := f1 + "," + f2
			origin := latitude + "," + longitude
			photoURL := getPhoto(nearby.Results[i].Photos[0].Photo_reference)
			content := "地址: " + nearby.Results[i].Vicinity + "\n" + countDistance(origin, destination)
			temp := linebot.NewCarouselColumn(
				"",
				nearby.Results[i].Name,
				content,
				linebot.NewURITemplateAction("查看街景", photoURL),
				linebot.NewURITemplateAction("開始導航", "http://maps.google.com/?q="+destination+""))

			s = append(s, temp)
		}
		template := linebot.NewCarouselTemplate(s...)
		templateMsg = linebot.NewTemplateMessage("Find Nearby branch", template)
	}
	return
}

func getJSON(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func getPhoto(ref string) (url string) {
	APIKey := os.Getenv("GoogleMapNearbySearchKey")
	maxwidth := "400"
	url = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=" + maxwidth + "&photoreference=" + ref + "&key=" + APIKey + ""
	return
}

func countDistance(origin, destination string) (output string) {
	APIkey := "AIzaSyAECvkl4TXtH9mXLJM-JvZ6LP6brfhYFDY"
	url := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=" + origin + "&destinations=" + destination + "&language=zh-TW&key=" + APIkey + ""
	log.Print(url)
	type Distance struct {
		Text  string `json:"text"`
		Value int    `json:"value"`
	}

	type Element struct {
		Distance Distance `json:"distance"`
	}

	type Row struct {
		Elements []Element `json:"elements"`
	}

	type DistanceResult struct {
		Rows   []Row  `json:"rows"`
		Status string `json:"status"`
	}

	distance := new(DistanceResult)
	err := getJSON(url, distance)
	if err != nil {
		panic(err)
	}

	if distance.Status == "OK" {
		output = "大約距離:" + distance.Rows[0].Elements[0].Distance.Text
	}

	return
}
