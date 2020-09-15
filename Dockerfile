FROM golang:1.15-alpine3.12 as builder

WORKDIR /go/src/app

ENV GOPATH=

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build

FROM alpine:3.12

COPY --from=builder /go/src/app/ls-proto-deps .
