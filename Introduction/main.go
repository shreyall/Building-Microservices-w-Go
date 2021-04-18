package main

import (
	"log"
	"microservices/introduction/handlers"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
