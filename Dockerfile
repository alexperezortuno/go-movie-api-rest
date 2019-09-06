ARG GO_VERSION=1.13
FROM golang:${GO_VERSION}-alpine
LABEL Description="Testing api movie" Vendor="Alex Peréz Ortuño" Version="1.0" maintainer="Alex Peréz <alexperezortuno@gmail.com>"
COPY src/api /go/src/api/
COPY src/main.go /go/src/
COPY installer.sh /go/
WORKDIR /go/src
EXPOSE 8090

RUN apk update && apk upgrade && apk add --no-cache bash git openssh ca-certificates
RUN chmod +x /go/installer.sh
RUN /bin/bash /go/installer.sh
ENTRYPOINT /go/src/api_movie
