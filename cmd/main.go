package main

import (
	"assignment1"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//start timer
	assignment1.InitTime()
	//trying to get portnumber
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", assignment1.HandlerNil)
	http.HandleFunc("/conservation/v1/country/", assignment1.CountryHandler)
	http.HandleFunc("/conservation/v1/species/", assignment1.SpeciesHandler)
	http.HandleFunc("/conservation/v1/diag/", assignment1.DiagHandler)
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
