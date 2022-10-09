FROM golang:1.19.2 AS builder

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /www

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o iot-echo

FROM alpine

LABEL author=github.com/wolanx
ENV TZ utc-8

RUN apk add --no-cache lua

WORKDIR /www

COPY --from=builder /www/iot-echo /usr/local/bin/
COPY config /root/.iot-echo

ENTRYPOINT ["iot-echo", "run"]

# docker build -f Dockerfile -t wolanx/iot-echo .
# docker run --restart=unless-stopped --name iotecho -d wolanx/iot-echo
