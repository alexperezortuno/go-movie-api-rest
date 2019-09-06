package v1

import (
	"api/v1/types"
	"encoding/json"
	"log"
	"net/http"
)

func popular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		errorResponse(405, "Method is not allowed", w)
		return
	}

	var page []string
	params := make([]types.Param, 2)

	page, pageOk := r.URL.Query()["page"]

	// Check if page is setted if not set new
	if !pageOk {
		page = append(page, "1")
		log.Println("Page not provided")
	}

	// Extra params for request
	pageParam := types.Param{Id: "page", Value: page[0]}
	langParam := types.Param{Id: "language", Value: "en-US"}

	params = append(params, pageParam)
	params = append(params, langParam)

	// Getting URL string
	endpoint := getUrl("movie/popular", params)

	// Request to movie popular endpoint
	req, detail := reqGet(w, endpoint)

	if req == nil {
		errorResponse(detail.Code, detail.Message, w)
		return
	}

	var responseObj types.PopularMovies
	_ = json.Unmarshal(req, &responseObj)
	_ = json.NewEncoder(w).Encode(responseObj)

	return
}

func movieDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := make([]types.Param, 1)

	id, ok := r.URL.Query()["id"]

	if !ok {
		errorResponse(500, "Service unavailable", w)
	}

	idParam := types.Param{Id: "id", Value: id[0]}
	params = append(params, idParam)
	endpoint := getUrl("movie/"+id[0], params)

	req, detail := reqGet(w, endpoint)

	if req == nil {
		errorResponse(detail.Code, detail.Message, w)
		return
	}

	var responseObj types.MovieDetail
	_ = json.Unmarshal(req, &responseObj)
	_ = json.NewEncoder(w).Encode(responseObj)
}
