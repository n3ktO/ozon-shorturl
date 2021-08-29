FROM golang:1.16-alpine AS builder
WORKDIR /go/src/shorturl
COPY . .
RUN go mod download
RUN go build -o ./build/shorturl ./src/main.go

FROM alpine:latest
COPY --from=builder /go/src/shorturl/build/shorturl /shorturl
ENTRYPOINT ["/shorturl"]