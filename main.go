package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		defer r.Body.Close()
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "hey %s\n", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
