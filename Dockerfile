# syntax=docker/dockerfile:1
##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

# Copy files needed to download dependencies
COPY go.mod ./
COPY go.sum ./

# Get dependencies
RUN go mod download

# Copy go files
COPY main.go .

COPY static static

# Build go executable
RUN go build -o /server

ENTRYPOINT ["/server"]