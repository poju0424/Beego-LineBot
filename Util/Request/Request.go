package Request

import (
	"H60Linebot/Util/File"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func ReuqestToServerWithRawBody(mainURL string, requestType string, resource string, mapHeader map[string]string, reqBody io.Reader) ([]byte, http.Header, error) {
	urlStr := mainURL

	requestURLInstance, _ := url.ParseRequestURI(urlStr)
	if resource != "" {
		requestURLInstance.Path = resource
		urlStr = fmt.Sprintf("%v", requestURLInstance)
	}

	client := &http.Client{}
	requestInstance, _ := http.NewRequest(requestType, urlStr, reqBody)

	for key, val := range mapHeader {
		requestInstance.Header.Add(key, val)
	}

	resp, e := client.Do(requestInstance)
	if e != nil {
		log.Println(e.Error())
		panic(e)
	}

	respBody, e := ioutil.ReadAll(resp.Body)

	log.Println("url " + urlStr)
	log.Println("requerst " + string(respBody))
	defer resp.Body.Close()
	return respBody, resp.Header, e
}

func ReuqestToServer(mainURL string, requestType string, resource string, mapHeader map[string]string, mapData map[string]string) ([]byte, http.Header, error) {

	requestData := url.Values{}
	for key, val := range mapData {
		requestData.Add(key, val)
	}
	encodedData := requestData.Encode()
	reqBody := bytes.NewBufferString(encodedData)

	mapHeader["Content-Length"] = strconv.Itoa(len(encodedData))

	return ReuqestToServerWithRawBody(mainURL, requestType, resource, mapHeader, reqBody)
}

func RequestUploadFile(uri string, mapHeader map[string]string, filePath string) ([]byte, http.Header, error) {
	body, contentType := File.ReadFile(filePath)
	mapHeader["Content-Type"] = contentType
	respBody, respHeader, err := ReuqestToServerWithRawBody(uri, "POST", "", mapHeader, body)
	return respBody, respHeader, err
}
