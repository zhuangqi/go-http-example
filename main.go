package main

import (
	"go-http-example/app"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", app.LoadRoute())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
