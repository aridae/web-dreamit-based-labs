FROM golang:1.15 as build
WORKDIR /project
COPY go.mod .
RUN go mod download
COPY . /project
RUN go build cmd/api_server/main.go

FROM ubuntu:latest as api-server
RUN apt update && apt install ca-certificates -y && rm -rf /var/cache/apt/*
COPY --from=build /project/main /
CMD ["./main"]