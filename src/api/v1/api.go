package v1

import (
    "gopkg.in/natefinch/lumberjack.v2"
    "log"
    "net/http"
    "os"
    "time"
)

func init() {
    // Server config
    s := &http.Server{
        Addr:           ":8090",
        Handler:        router(),
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    // Configure Logging
    LogFileLocation := os.Getenv("LOG_FILE_LOCATION")

    if LogFileLocation != "" {
        log.SetOutput(&lumberjack.Logger{
            Filename:   LogFileLocation,
            MaxSize:    500, // megabytes
            MaxBackups: 3,
            MaxAge:     5,   //days
            Compress:   true, // disabled by default
        })
    }

    log.Println("Starting Server, Version: ", version)

    if err := s.ListenAndServe(); err != nil {
        log.Fatal(err)
    }

    waitForShutdown(s)
}
