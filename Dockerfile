FROM golang:1.15 as builder
WORKDIR /usr/src/build
COPY main.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o knights_on_a_keypad

FROM alpine:3
WORKDIR /usr/bin
COPY --from=builder /usr/src/build/knights_on_a_keypad .
CMD ["/usr/bin/knights_on_a_keypad"]
