FROM golang:1.11.4 AS builder
RUN go version

COPY . /go/src/gitee.com/odd-socket
WORKDIR /go/src/gitee.com/odd-socket

RUN set -x && \
    go get -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/

COPY --from=builder /go/src/gitee.com/odd-socket/app .
COPY ./config/config.yml /root/config/config.yml

EXPOSE 8000
ENTRYPOINT ["./app"]
