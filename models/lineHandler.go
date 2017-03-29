package models

import (
	"encoding/json"
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
				msg, isValid := spliteTextMsg(message.Text)
				if isValid {
					if _, err = bot.ReplyMessage(event.ReplyToken, ReplyTemplateMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				}
			case *linebot.LocationMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, getNerybyBank(message.Latitude, message.Longitude)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.VideoMessage:
				log.Print(message.ID)
				log.Print(message.OriginalContentURL)
				log.Print(message.PreviewImageURL)
			case *linebot.ImageMessage:
				log.Print(message.ID)
				log.Print(message.OriginalContentURL)
				log.Print(message.PreviewImageURL)
			case *linebot.AudioMessage:
				log.Print(message.ID)
				log.Print(message.OriginalContentURL)
				log.Print(message.Duration)

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
	log.Print(lat)
	log.Print(lon)
	latitude := strconv.FormatFloat(lat, 'f', -1, 64)
	longitude := strconv.FormatFloat(lon, 'f', -1, 64)
	name := "臺灣銀行股份有限公司"
	APIKey := os.Getenv("GoogleMapNearbySearchKey")
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + latitude + "," + longitude + "&name=" + name + "&key=" + APIKey + "&language=zh-TW&types=bank&rankby=distance"

	type LatLng struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
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
	getJSON(url, nearby)

	var s []*linebot.CarouselColumn
	if nearby.Status == "OK" {
		for i := 0; i < 5; i++ {
			loc := nearby.Results[i].Geometry.Location.Lat + "," + nearby.Results[i].Geometry.Location.Lng
			photoURL := getPhoto(nearby.Results[i].Photos[0].Photo_reference)
			log.Print(photoURL)
			temp := linebot.NewCarouselColumn(
				photoURL,
				nearby.Results[i].Name, nearby.Results[i].Vicinity,
				linebot.NewURITemplateAction("開始導航", "http://maps.google.com/?q="+loc+""))
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
