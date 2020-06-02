package main

import (
	"net/http"
	"log"
	"homelibrary/internal/handles"
	"homelibrary/config"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",handles.HandleRequest)

	log.Fatal(http.ListenAndServe(config.Port,mux))
}
