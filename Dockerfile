# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13.1

WORKDIR /go/src/github.com/trinhquocviet/go-echo-skeleton
