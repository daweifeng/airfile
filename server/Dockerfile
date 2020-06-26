FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/airfile-server
COPY ./src/ .
RUN go get ./...
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/airfile-server ./main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/airfile-server/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/airfile-server --port 8080

