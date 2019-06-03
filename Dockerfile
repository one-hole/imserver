FROM golang:1.12.1 AS builder
RUN go version

COPY . /go/src/github.com/one-hole/imserver
WORKDIR /go/src/github.com/one-hole/imserver

RUN set -x && \
    go get -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/

COPY --from=builder /go/src/github.com/one-hole/imserver/app .
COPY ./config/config.yml /root/config/config.yml

EXPOSE 8000
ENTRYPOINT ["./app"]
