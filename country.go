package assignment1

import (
	"encoding/json"
	"net/http"
	"strings"
)

//CountryHandler function
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	APIURL := "https://restcountries.eu/rest/v2/alpha/"
	//gets the query from the url the user writes
	urlQuery := r.URL.RawQuery
	//splitting the url path that the user writes
	parts := strings.Split(r.URL.Path, "/")

	//checking if the path is right
	if len(parts) != 5 || parts[1] != "conservation" || parts[2] != "v1" || parts[3] != "country" || parts[4] == "" {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format /conservation/v1/country/AlphaCode?limit='integer'", status)
		return
	}
	APIURL += parts[4]

	//setting up a client and a empty country strct
	client := http.DefaultClient
	country := &Country{}

	//do a request on the APIURL and decode the response to the countrystruct
	resp := GetTheRequest(APIURL, client)
	DecodeCountry(country, resp, w)

	//creating a new URL to get species and species keys
	NEWURL := "http://api.gbif.org/v1/occurrence/search?country=" + country.Alpha2Code + "&" + urlQuery

	//do a request on the new url and decodes the response into the countrystruct
	resp2 := GetTheRequest(NEWURL, client)
	DecodeKeyAndSpecies(country, resp2)
	//encode countrystruct back to the user
	json.NewEncoder(w).Encode(country)

}

//DecodeCountry function
func DecodeCountry(country *Country, resp *http.Response, w http.ResponseWriter) {

	//checking if there is any errors from the decoding
	err := json.NewDecoder(resp.Body).Decode(&country)
	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, "Expected a country, got nothing", status)
		return
	}
}

//DecodeKeyAndSpecies function
func DecodeKeyAndSpecies(country *Country, resp *http.Response) {

	//creating a empty resultstruct and decodes it
	speciesKeyAndSpecies := &Results{}
	json.NewDecoder(resp.Body).Decode(speciesKeyAndSpecies)

	var species []string
	var spesiesKey []int

	//loop that puts all the decoded values into two arrays that contains keys and speciesname
	for _, v := range speciesKeyAndSpecies.Result {
		species = append(species, v.Specie)
		spesiesKey = append(spesiesKey, v.Key)

	}

	//copy keys and speciesname into country
	country.SpeciesKey = spesiesKey
	country.Species = species

}
