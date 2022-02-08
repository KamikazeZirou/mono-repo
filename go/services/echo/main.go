package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	body := r.Form.Get("body")
	_, _ = fmt.Fprint(w, body)
}
