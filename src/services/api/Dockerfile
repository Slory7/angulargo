FROM golang:alpine as builder
ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED=0

RUN mkdir /build 
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./src ./src

WORKDIR /build/src/services/api

RUN go build -ldflags "-s -w" -o main .
FROM alpine
RUN apk add -U --no-cache ca-certificates && rm -rf /var/cache/apk/* /tmp/* 
COPY --from=builder /build/src/services/api/main /app/
COPY --from=builder /build/src/services/api/config.json /app/

WORKDIR /app
EXPOSE 8010
ENTRYPOINT ["./main"]