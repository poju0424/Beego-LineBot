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
)

func GetRateInfo(request string) (message string) {
	body, header := ReadFile()
	datetime := GetTimeFromFileName(header)

	r := bytes.NewReader(body)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		matched, err := regexp.MatchString("^("+request+")", line)
		Debug.CheckErr(err)
		if matched {
			arr := strings.Split(line, ",")
			message = "台銀" + arr[0] + "即時匯率:" +
				"\n 現金買入:" + arr[2] +
				"\n 現金賣出:" + arr[3] +
				"\n 即期買入:" + arr[12] +
				"\n 即期賣出:" + arr[13] +
				"\n 更新時間(" + datetime + ")"
		}
	}
	if len(message) <= 0 {
		message = ""
	}
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
