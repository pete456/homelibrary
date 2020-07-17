package main

import (
	"net/http"
	"log"
	"os"
	"homelibrary/config"
	"homelibrary/internal/daemon"
	"homelibrary/internal/handles"
)

func client(args []string) {
	if len(args) == 1 {
		daemon.StartDaemon([]uintptr{})
	}
	if
	os.Exit(0)

}

func main() {
	if os.Getppid() != 1 {
		client(os.Args)
	}

	lf,err := os.Create(config.LogFile)
	defer lf.Close()
	if err != nil {
		log.Fatalf("Error openning log %s",err)
		os.Exit(1)
	}
	log.SetOutput(lf)

	mux := http.NewServeMux()
	mux.HandleFunc("/",handles.HandleRequest)

	log.Fatal(http.ListenAndServe(config.Port,mux))
}
