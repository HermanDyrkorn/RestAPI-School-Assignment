package assignment1

import (
	"net/http"
)

//GetTheRequest function returns a respons from the api
func GetTheRequest(URL string, client *http.Client) *http.Response {
	//do a request on the url
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		panic(err)
	}
	//getting response from the client
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//returning the response
	return resp
}

//HandlerNil function that deals with the / endpoint
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid request, expected formating conservation/v1/diag or country or species", http.StatusBadRequest)
}
