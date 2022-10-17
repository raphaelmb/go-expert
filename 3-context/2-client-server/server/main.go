package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request completed")
	select {
	case <-time.After(5 * time.Second):
		// logs on server stdout
		log.Println("Request successfully processed")
		// logs on browser
		w.Write([]byte("Request successfully processed"))
	case <-ctx.Done():
		// logs on server stdout
		log.Println("Request cancelled by client")
	}
}
