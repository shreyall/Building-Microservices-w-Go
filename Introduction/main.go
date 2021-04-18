package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World! ")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "oooppps", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hi there, %s", d)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
