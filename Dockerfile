FROM golang:1.24.2-alpine
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/command
WORKDIR /go/src/command
ADD . /go/src/command