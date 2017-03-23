package models

import (
	"bufio"
	"bytes"
	"fmt"
	"hello/Util/Debug"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GetRateInfo(request string) (message string) {
	body, header := ReadFile()
	datetime := GetTimeFromFileName(header)

	r := bytes.NewReader(body)
	scanner := bufio.NewScanner(r)
	var currency, cashBuy, cashSell, rateBuy, rateSell string
	for scanner.Scan() {
		line := scanner.Text()
		matched, err := regexp.MatchString("("+request+")", line)
		Debug.CheckErr(err)
		fmt.Print(request)
		if matched {
			arr := strings.Split(line, ",")
			currency = arr[0]
			cashBuy = arr[2]
			cashSell = arr[3]
			rateBuy = arr[12]
			rateSell = arr[13]
		} else {
			return "404"
		}

	}
	message = "台銀" + currency + "即時匯率:" +
		"\n 現金買入:" + cashBuy +
		"\n 現金賣出:" + cashSell +
		"\n 即期買入:" + rateBuy +
		"\n 即期賣出:" + rateSell +
		"\n 更新時間(" + datetime + ")"
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

func ConnectDB(currency string, cashbuy, cashsell, ratebuy, ratesell float64) {

}
