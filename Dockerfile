# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-alpine3.15 AS build

WORKDIR /restapi

COPY . .
COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go build -o restapi ./cmd/main.go

# Deploy
FROM alpine:3.15

WORKDIR /restapi

COPY --from=build /restapi /restapi

CMD [ "./restapi" ]
