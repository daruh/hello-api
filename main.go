package main

import (
	"hello-api/handlers/rest"
	"log"
	"net/http"
)

func main() {

	addr := ":8083"

	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)
	log.Println("config map demo")

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct { // <6>
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
