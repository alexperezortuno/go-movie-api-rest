package main

import (
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

type Results struct {
	Popularity  float32 `json:"popularity"`
	VoteCount   int16	`json:"vote_count"`
	Video       bool    `json:"video"`
	PosterPath  string  `json:"poster_path"`
	Title       string  `json:"title"`
	VoteAverage float32 `json:"vote_average"`
	Id          int32   `json:"id"`
	Overview	string	`json:"overview"`
}

type PopularResponse struct {
	Page 			int 		`json:"page"`
	TotalResults 	int16 		`json:"total_results"`
	TotalPages 		int 		`json:"total_pages"`
	Results			[]Results	`json:"results"`
}

type Param struct {
	id 		string
	value 	string
}

type Error struct {
	Code	int16	`json:"code"`
	Message string	`json:"Message"`
}

func getUrl(endpoint string, params []Param) string {
	var timb_url bytes.Buffer
	v := url.Values{}

	v.Set("api_key", os.Getenv("API_KEY"))

	for i := 0; i < len(params); i++  {
		v.Add(params[i].id, params[i].value)
	}

	timb_url.WriteString(os.Getenv("TIMDB_URL"))
	timb_url.WriteString(endpoint)
	timb_url.WriteString("?")
	timb_url.WriteString(v.Encode())

	return timb_url.String()
}

func getPopular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var page []string
	var params []Param

	page, pageOk := r.URL.Query()["page"]

	// Check if page is setted if not set new
	if !pageOk {
		page = append(page, "1")
		log.Println("Page not provided")
	}

	// Extra params for request
	pageParam := Param{id: "page", value: page[0]}
	langParam := Param{id: "language", value: "en-US"}

	params = append(params, pageParam)
	params = append(params, langParam)

	// Getting URL string
	endpoint := getUrl("popular", params)

	// Request to movie popular endpoint

	response, err := http.Get(endpoint)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		_ = json.NewEncoder(w).Encode(&Error{Code:500, Message: "Service not available"})
	}

	if response.StatusCode == 200 {
		data, dataErr := ioutil.ReadAll(response.Body)

		if dataErr != nil {
			log.Fatal(dataErr)
		}

		var responseObj PopularResponse
		_ = json.Unmarshal(data, &responseObj)
		_ = json.NewEncoder(w).Encode(responseObj)
	} else {
		log.Printf("The HTTP request failed with error %s\n", err)
		_ = json.NewEncoder(w).Encode(&Error{Code: int16(response.StatusCode), Message: "Service not available"})
	}
}

func getVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	fmt.Println("Starting app ...")

	r := mux.NewRouter()
	r.HandleFunc("/popular", getPopular).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", r))
}