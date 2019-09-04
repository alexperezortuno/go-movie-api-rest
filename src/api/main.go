package main

import (
	"api/interfaces"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func errorResponse(code int16, message string, w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(&interfaces.Error{Code: code, Message: message})
}

func getUrl(endpoint string, params []interfaces.Param) string {
	var timb_url bytes.Buffer
	v := url.Values{}

	v.Set("api_key", os.Getenv("API_KEY"))

	for i := 0; i < len(params); i++ {
		v.Add(params[i].Id, params[i].Value)
	}

	timb_url.WriteString(os.Getenv("TIMDB_URL"))
	timb_url.WriteString(endpoint)
	timb_url.WriteString("?")
	timb_url.WriteString(v.Encode())

	return timb_url.String()
}

func reqGet(w http.ResponseWriter, endpoint string) []byte {
	req, err := http.Get(endpoint)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		errorResponse(500, "Service not available", w)
	}

	if req.StatusCode == 200 {
		data, dataErr := ioutil.ReadAll(req.Body)

		if dataErr != nil {
			log.Fatal(dataErr)
		}

		return data
	} else {
		log.Printf("The HTTP request failed with error %s \n", err)
		// errorResponse(int16(req.StatusCode), "Service not available", w)
	}

	return nil
}

func popular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var page []string
	var params []interfaces.Param

	page, pageOk := r.URL.Query()["page"]

	// Check if page is setted if not set new
	if !pageOk {
		page = append(page, "1")
		log.Println("Page not provided")
	}

	// Extra params for request
	pageParam := interfaces.Param{Id: "page", Value: page[0]}
	langParam := interfaces.Param{Id: "language", Value: "en-US"}

	params = append(params, pageParam)
	params = append(params, langParam)

	// Getting URL string
	endpoint := getUrl("movie/popular", params)

	// Request to movie popular endpoint
	reqGet(w, endpoint)
}

func movieDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params []interfaces.Param

	id, ok := r.URL.Query()["id"]

	if !ok {
		errorResponse(500, "Service unavailable", w)
	}

	idParam := interfaces.Param{Id: "id", Value: id[0]}
	params = append(params, idParam)
	endpoint := getUrl("movie/"+id[0], params)

	req := reqGet(w, endpoint)

	if req == nil {
		errorResponse(500, "Service not available", w)
		return
	}

	var responseObj interfaces.MovieDetail
	_ = json.Unmarshal(req, &responseObj)
	_ = json.NewEncoder(w).Encode(responseObj)
}

func main() {
	fmt.Println("Starting app, version ", version)

	r := mux.NewRouter()
	r.HandleFunc("/popular", popular).Methods("GET")
	r.HandleFunc("/movie", movieDetail).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", r))
}
