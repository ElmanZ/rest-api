# syntax=docker/dockerfile:1
##Build
#Alpine chosen because of its relatively small size
FROM golang:1.16-alpine3.15 AS build

WORKDIR /restapi

COPY . .

#Download neccessary go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

#Compile the application
RUN go build -o restapi ./cmd/main.go

##Deploy
FROM alpine:3.15

WORKDIR /restapi

COPY --from=build /restapi /restapi

#Tells Docker what command to execute when the image is used to start a container
CMD [ "./restapi" ]
