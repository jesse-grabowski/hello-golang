# Learning GO Webapp

## Prerequisites

pkger CLI

    go install github.com/markbates/pkger/cmd/pkger

## Build & Run

Generate pkger data (embedded files like application.yml)

    go generate

Run application using port 8080 and local profile

    go run . /port 8080 /profile local

## Docker

Build container image

    docker build -t test-go-build . 

Run container image

    docker run -p 8080:8080 test-go-build -p 8080 -P local -l /go/bin/
