# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as build-env

RUN apk add --update --no-cache ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/hello

FROM scratch
COPY --from=build-env /go/bin/hello /go/bin/hello
COPY application-local.yml /go/bin/
ENTRYPOINT ["/go/bin/hello"]