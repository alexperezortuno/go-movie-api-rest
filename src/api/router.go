package api

import (
    "github.com/gorilla/mux"
)

func router() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/popular", popular).Methods("GET")
    r.HandleFunc("/movie", movieDetail).Methods("GET")
    
    return r
}
