# Learning GO Webapp

## Prerequisites

pkger CLI

    go install github.com/markbates/pkger/cmd/pkger

## Build & Run

Generate pkger data (embedded files like application.yml)

    go generate

Run application using port 8080 and local profile

    go run . /port 8080 /profile local