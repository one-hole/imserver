FROM golang:1.12.1 AS builder
RUN go version

COPY . /go/src/github.com/w-zengtao/socket-server
WORKDIR /go/src/github.com/w-zengtao/socket-server

RUN set -x && \
    go get -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/

COPY --from=builder /go/src/github.com/w-zengtao/socket-server/app .
COPY ./config/config.yml /root/config/config.yml

EXPOSE 8000
ENTRYPOINT ["./app"]
