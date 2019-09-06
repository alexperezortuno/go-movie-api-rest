#!/bin/bash
cd ${GOPATH}/src
go get -t gopkg.in/natefinch/lumberjack.v2
go get -t github.com/gorilla/mux
go build -o api_movie
cd ${GOPATH}
mkdir -p ${GOPATH}/logs/api_movie
chmod 755 ${GOPATH}/logs/api_movie