package controllers

import (
	"log"
	"net/http"
)

type PushHandler struct{}

func (*PushHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Body)
	log.Print(r.RequestURI)
	// params := strings.Split(r.RequestURI, "/")
	// if len(params) == 4 {
	// 	time := params[2]
	// 	name := params[3]
	// 	data := model.GetData(time, name)
	// 	buff := createChart(data)
	// 	w.Header().Set("Content-Type", "image/jpeg")
	// 	w.Header().Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
	// 	if _, err := w.Write(buff.Bytes()); err != nil {
	// 		log.Print(err)
	// 	}
	return
	// }

}
