FROM golang:1.12.5 AS builder
RUN go version

COPY . /imserver
WORKDIR /imserver

# RUN set -x && \

RUN export GOPROXY=https://goproxy.io && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/

COPY --from=builder /imserver/app .
# COPY ./config/config.yml /root/config/config.yml

EXPOSE 8000
ENTRYPOINT ["./app"]
