package assignment1

import (
	"encoding/json"
	"net/http"
	"strings"
)

//DiagHandler function
func DiagHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	//splitting the path into smaller parts
	parts := strings.Split(r.URL.Path, "/")
	//checks that the path is right
	if len(parts) != 5 || parts[1] != "conservation" || parts[2] != "v1" || parts[3] != "diag" {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format /conservation/v1/diag/", status)
		return
	}
	//creating a diagnostics struct
	diag := &Diag{}

	//filling the diagstruct with info
	diag.Uptime = Uptime()
	diag.Version = "V1"
	diag.RestCountry = GetStatusCode("https://restcountries.eu")
	diag.GBIF = GetStatusCode("http://api.gbif.org/v1/")
	//encode the struct back to the user
	json.NewEncoder(w).Encode(diag)

}

//GetStatusCode function
func GetStatusCode(URL string) int {
	//getrequest on a given url
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	//returning the statuscode of the response
	return resp.StatusCode
}
