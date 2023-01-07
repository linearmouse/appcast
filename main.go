package main

import (
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/appcast.xml" {
		http.NotFound(w, r)
		return
	}

	appcast, err := getAppCast()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/xml")
	w.Write(appcast)
}

func main() {
	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
