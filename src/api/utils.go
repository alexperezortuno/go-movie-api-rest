package api

import (
	"api/types"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func errorResponse(code int16, message string, w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(&types.Error{Code: code, Message: message})
	return
}

func reqGet(w http.ResponseWriter, endpoint string) ([]byte, types.ResponseDetail) {
	req, err := http.Get(endpoint)

	if err != nil {
		log.Printf("Error %s \n", err)
		return nil, types.ResponseDetail{Code: 1100, Message: "Bad request"}
	}

	if req.StatusCode == 200 {
		data, dataErr := ioutil.ReadAll(req.Body)

		if dataErr != nil {
			log.Fatal(dataErr)
		}

		return data, types.ResponseDetail{Code: 1200, Message: "OK"}
	} else {
		log.Printf("The HTTP request failed with error %s \n", err)
	}

	return nil, types.ResponseDetail{Code: 1300, Message: "Service unavailable"}
}

func getUrl(endpoint string, params []types.Param) string {
	var timdbUrl bytes.Buffer
	v := url.Values{}

	v.Set("api_key", os.Getenv("API_KEY"))

	for i := 0; i < len(params); i++ {
		v.Add(params[i].Id, params[i].Value)
	}

	timdbUrl.WriteString(os.Getenv("TIMDB_URL"))
	timdbUrl.WriteString(endpoint)
	timdbUrl.WriteString("?")
	timdbUrl.WriteString(v.Encode())

	return timdbUrl.String()
}
