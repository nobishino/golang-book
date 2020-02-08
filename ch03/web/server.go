package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/nobishino/golang-book/ch03/svg"
)

func Listen(port int) {
	http.HandleFunc("/", handler)
	host := fmt.Sprintf("localhost:%d", port)
	log.Fatal(http.ListenAndServe(host, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		log.Println(k, v)
	}
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	s := svg.Setting{}
	for k, v := range r.Form {
		log.Println(k, v)
		setParam(k, v[0], &s)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	svg.WriteWithSetting(w, s)
}

//パラメータのセッティングに成功したらtrueを返す
func setParam(key, value string, s *svg.Setting) bool {
	v, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	switch key {
	case "w":
		s.Width = v
		return true
	case "h":
		s.Height = v
		return true
	case "r":
		s.Rgb.Red = v
		return true
	case "g":
		s.Rgb.Green = v
		return true
	case "b":
		s.Rgb.Blue = v
		return true
	case "c":
		s.Rgb.Change = v > 0
		return true
	}
	return false
}
