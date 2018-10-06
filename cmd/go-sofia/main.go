package main

import "log"
import "fmt"
import "net/http"
import "os"

import "github.com/gorilla/mux"
import "github.com/ialidzhikov/go-sofia/internal/diagnostics"

func main() {
	log.Print("Starting the application...")

	blPort := os.Getenv("PORT")
	if len(blPort) == 0 {
		log.Fatal("The application port should be set")
	}

	diagnosticsPort := os.Getenv("DIAG_PORT")
	if len(diagnosticsPort) == 0 {
		log.Fatal("The diagnostics port should be set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func() {
		error := http.ListenAndServe(":" + blPort, router)
		if error != nil {
			log.Fatal(error)
		}
	}()

	diagnostics := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":" + diagnosticsPort, diagnostics)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}