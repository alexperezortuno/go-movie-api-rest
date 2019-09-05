package api

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func init() {
    fmt.Println("Starting app, version ", version)
    
    s := &http.Server{
        Addr:           ":8090",
        Handler:        router(),
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    
    log.Fatal(s.ListenAndServe())
}
