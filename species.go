package assignment1

import (
	"encoding/json"
	"net/http"
	"strings"
)

//SpeciesHandler function
func SpeciesHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	APIURL := "http://api.gbif.org/v1/species/"

	//splitting the path into smaller parts
	parts := strings.Split(r.URL.Path, "/")

	//checks that the path is right
	if len(parts) != 5 || parts[1] != "conservation" || parts[2] != "v1" || parts[3] != "species" || parts[4] == "" {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format /conservation/v1/species/'Spiciesnumber'", status)
		return
	}
	APIURL += parts[4]

	//setting up a client and does a response
	client := http.DefaultClient
	resp := GetTheRequest(APIURL, client)

	//appending "name" to the APRURL and does a new response
	APIURL += "/name"
	resp2 := GetTheRequest(APIURL, client)

	//creating a species struct and a speciesyear struct
	species := &Species{}
	year := &SpeciesYear{}

	//decode species and speciesyear
	DecodeSpecies(species, resp, w)
	DecodeYear(resp2, year)

	//copy over year into speciesstruct and encode it back to user
	species.Year = year.Year
	json.NewEncoder(w).Encode(species)

}

//DecodeSpecies function
func DecodeSpecies(species *Species, resp *http.Response, w http.ResponseWriter) {
	//checking if there is any errors from the decoding
	err := json.NewDecoder(resp.Body).Decode(species)
	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, "Expected a species, got nothing", status)
		return
	}

}

//DecodeYear function
func DecodeYear(resp *http.Response, year *SpeciesYear) {
	json.NewDecoder(resp.Body).Decode(year)
}
