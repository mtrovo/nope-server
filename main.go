package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(`<html><h1>NOPE</h1></html>`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	log.Println("server running at", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("listen error:", err)
	}

}
