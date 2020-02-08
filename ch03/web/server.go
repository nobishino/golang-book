package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nobishino/golang-book/ch03/svg"
)

func Listen(port int) {
	http.HandleFunc("/", handler)
	host := fmt.Sprintf("localhost:%d", port)
	log.Fatal(http.ListenAndServe(host, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	svg.Write(w)
	// io.WriteString(w, "Hello SVG")
}
