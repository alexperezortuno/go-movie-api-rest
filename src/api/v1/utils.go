package v1

import (
	"api/v1/types"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
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
		log.Printf("Requesting to %s \n ", )
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

func waitForShutdown(s *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	s.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}