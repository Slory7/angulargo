FROM golang:alpine as builder
ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED=0

RUN mkdir /build 
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./src/ ./src

WORKDIR /build/src/services/api/cmd/web

RUN go build -ldflags "-s -w" -o main .

FROM alpine
RUN apk add -U --no-cache ca-certificates && rm -rf /var/cache/apk/* /tmp/* 
COPY --from=builder /build/src/services/api/cmd/web/main /app/
COPY --from=builder /build/src/services/api/cmd/web/config.json /app/

WORKDIR /app
EXPOSE 8011
ENTRYPOINT ["./main"]