package main

import "log"
import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import "github.com/ialidzhikov/go-sofia/internal/diagnostics"

func main() {
	log.Print("Hello, world")

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	router.HandleFunc("/health")

	go func() {
		error := http.ListenAndServe(":8080", router)
		if error != nil {
			log.Fatal(error)
		}
	}()

	diagnostics := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":8585", diagnostics)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}