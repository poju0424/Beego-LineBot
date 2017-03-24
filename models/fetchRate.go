package models

import (
	"bufio"
	"bytes"
	"hello/Util/Debug"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func getRateInfo(request string) (message []string) {
	body, header := ReadFile()
	datetime := GetTimeFromFileName(header)
	code, name := fuzzySearch(request)

	r := bytes.NewReader(body)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		matched, err := regexp.MatchString("^("+code+")", line)
		Debug.CheckErr(err)
		if matched {
			arr := strings.Split(line, ",")
			// message = "台銀" + name + "即時匯率:" +
			// 	"\n 現金買入:" + arr[2] +
			// 	"\n 現金賣出:" + arr[3] +
			// 	"\n 即期買入:" + arr[12] +
			// 	"\n 即期賣出:" + arr[13] +
			// 	"\n 更新時間(" + datetime + ")"
			message[0] = "台銀" + name + "即時匯率:"
			message[1] = "\n 現金買入:" + arr[2] +
				"\n 現金賣出:" + arr[3] +
				"\n 即期買入:" + arr[12] +
				"\n 即期賣出:" + arr[13] +
				"\n 更新時間(" + datetime + ")"
			message[2] = name
		}
	}
	if len(message) <= 0 {
		return nil
	}
	return
}

func ReplyTemplateMessage(request string) (templateMsg linebot.Message) {
	var AltText = "alttext"
	// var template linebot.Template
	msg := getRateInfo(request)
	template := linebot.NewButtonsTemplate(
		"", msg[0], msg[1],
		linebot.NewURITemplateAction("Go to Taiwan Bank Website", "http://rate.bot.com.tw/xrt?Lang=zh-TW"),
		linebot.NewPostbackTemplateAction("Find nearby branch", "回傳SERVER值", "不跟你說"),
		linebot.NewMessageTemplateAction("Query rate again", msg[2]),
	)

	templateMsg = linebot.NewTemplateMessage(AltText, template)
	return
}

func ReadFile() (body []byte, header http.Header) {
	var filePath = "http://rate.bot.com.tw/xrt/flcsv/0/day"
	resp, err := http.Get(filePath)
	Debug.CheckErr(err)

	header = resp.Header
	respBody, e := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e == nil {
		body = respBody
	}
	return
}

func GetTimeFromFileName(header http.Header) (output string) {
	contentDispostion := header.Get("Content-Disposition")
	index := strings.Index(contentDispostion, "@") + 1
	datetime := contentDispostion[index : index+12]
	t, err := time.Parse("200601021504", datetime)
	Debug.CheckErr(err)
	output = t.Format("2006/01/02 15:04")
	return
}

func connectDB(currency string, cashbuy, cashsell, ratebuy, ratesell float64) {

}

func fuzzySearch(msg string) (code, name string) {
	searchList := [][]string{}
	searchList = append(searchList, []string{"日", "jpy", "日圓"})
	searchList = append(searchList, []string{"jp", "jpy", "日圓"})
	searchList = append(searchList, []string{"美", "usd", "美金"})
	searchList = append(searchList, []string{"us", "usd", "美金"})
	searchList = append(searchList, []string{"人民", "cny", "人民幣"})
	searchList = append(searchList, []string{"rmb", "cny", "人民幣"})
	searchList = append(searchList, []string{"cn", "cny", "人民幣"})
	searchList = append(searchList, []string{"歐", "eur", "歐元"})
	searchList = append(searchList, []string{"eu", "eur", "歐元"})
	searchList = append(searchList, []string{"港", "hkd", "港幣"})
	searchList = append(searchList, []string{"hk", "hkd", "港幣"})
	searchList = append(searchList, []string{"kr", "krw", "韓元"})
	searchList = append(searchList, []string{"韓", "krw", "韓元"})

	max := len(searchList)
	var found = false
	code = "404"
	name = "not found any match key word"
	for i := 0; i < max; i++ {
		if strings.Contains(strings.ToLower(msg), searchList[i][0]) {
			if found {
				code = "404"
				name = "more than one key word have found"
				return
			} else {
				found = true
			}
			code = strings.ToUpper(searchList[i][1])
			name = searchList[i][2]
		}
	}
	return
}
