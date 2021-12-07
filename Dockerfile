FROM golang:1.17

WORKDIR /src
RUN go install golang.org/dl/gotip@latest
RUN gotip download
