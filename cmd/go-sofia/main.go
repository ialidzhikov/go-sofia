package main

import "log"
import "net/http"

func main() {
	log.Print("Hello, world")

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Fatal(error)
	}

}