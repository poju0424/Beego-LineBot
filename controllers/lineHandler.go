package controllers

import (
	"encoding/json"
	"hello/service"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

type LineHandler struct{}

func (*LineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	// log.Println("Bot:", bot, " err:", err)

	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				log.Print(message.Text)
				msg := message.Text
				// msg, isValid := spliteTextMsg(message.Text)
				// if isValid {
				// 	if _, err = bot.ReplyMessage(event.ReplyToken, service.ReplyTemplateMessage(msg)).Do(); err != nil {
				// 		log.Print(err)
				// 	}
				// }
				if _, err = bot.ReplyMessage(event.ReplyToken, service.ReplyTemplateMessage(msg)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, getNerybyBank(message.Latitude, message.Longitude)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.AudioMessage:
				url := "https://beegolinebot.herokuapp.com/currency/ltm/JPY"
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(url, url)).Do(); err != nil {
					log.Print(err)
				}
			}
		} else if event.Type == linebot.EventTypePostback {
			text := event.Postback.Data
			byteArr := []byte(text)
			var postbackObj service.PostbackObj
			err := json.Unmarshal(byteArr, &postbackObj)
			if err != nil {
				log.Print(err)
			}
			switch postbackObj.Method {
			case "text":
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(postbackObj.Data)).Do()
			case "image":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(postbackObj.Data, postbackObj.Data)).Do(); err != nil {
					log.Print(err)
				}
			}

		}
	}
}

func spliteTextMsg(msg string) (subMsg string, isValid bool) {
	re := regexp.MustCompile("^&&(.*)")
	arr := re.FindStringSubmatch(msg)
	if len(arr) > 0 {
		isValid = true
		subMsg = arr[1]
	} else {
		isValid = false
	}
	return
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func getNerybyBank(lat, lon float64) (templateMsg linebot.Message) {
	latitude := strconv.FormatFloat(lat, 'f', -1, 64)
	longitude := strconv.FormatFloat(lon, 'f', -1, 64)
	name := "臺灣銀行股份有限公司"
	APIKey := os.Getenv("GoogleMapNearbySearchKey")
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + latitude + "," + longitude + "&name=" + name + "&key=" + APIKey + "&language=zh-TW&types=bank&rankby=distance"

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

func getPhoto(ref string) (url string) {
	APIKey := os.Getenv("GoogleMapNearbySearchKey")
	maxwidth := "400"
	url = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=" + maxwidth + "&photoreference=" + ref + "&key=" + APIKey + ""
	return
}

func countDistance(origin, destination string) (output string) {
	APIkey := "AIzaSyAECvkl4TXtH9mXLJM-JvZ6LP6brfhYFDY"
	url := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=" + origin + "&destinations=" + destination + "&language=zh-TW&key=" + APIkey + ""

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
