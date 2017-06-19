package controllers

import (
	"log"
	"net/http"
)

type SensorDataHandler struct{}

func (*SensorDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := "https://h60jwt.azurewebsites.net/V1/Device/CE87340D767A/SensorData/AirQuality/History/7"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("clientId", "test123")
	req.Header.Set("Authorization", "bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1bmlxdWVfbmFtZSI6InRlc3QxMjMiLCJ1c2VyTmFtZSI6ImZyYW5rIiwid2lmaW1hYyI6IiIsInVzZXJSb2xlIjoiMSIsInJvbGUiOiJVc2VyIiwidXNlcklkIjoiYWE4MzQ5M2YtMTE3My00MDkzLThmNzctMGMxYWQxM2E3ZDE2IiwiZW1haWwiOiJxbWl0dzE4OEBnbWFpbC5jb20iLCJpc3MiOiJodHRwczovL2g2MGp3dC5henVyZXdlYnNpdGVzLm5ldCIsImF1ZCI6IjBmMTMzZDFjMzQ3ZTQ0Zjg5ZmUxYmMwNWE2NDdhOGJhIiwiZXhwIjoxNTAwNDQzMzgwLCJuYmYiOjE0OTc4NTEzODB9.m3sJJzF8Vvz9KRb9xnPPsp1OARr72DekYfN4qRA6-YU")
	client := &http.Client{}
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Print(err)
	}
	log.Print(resp)
}
