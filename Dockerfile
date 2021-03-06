FROM golang:1.12.6 AS builder
RUN go version

COPY . /imserver
WORKDIR /imserver

RUN set -x && \
    export GOPROXY=https://goproxy.io && \
    go get -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o imserver .

FROM scratch
WORKDIR /root/

COPY --from=builder /imserver/imserver .

EXPOSE 8000
ENTRYPOINT ["./imserver"]
