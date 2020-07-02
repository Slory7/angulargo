FROM golang:alpine as builder
ENV GOPROXY https://goproxy.io
ENV CGO_ENABLED=0

RUN apk add -U --no-cache ca-certificates

RUN mkdir /build 
WORKDIR /build

ADD ./src ./src
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /build/src/services/gather

RUN go build -ldflags "-s -w" -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/src/services/gather/main /app/
#COPY --from=builder /build/src/services/gather/config.json /app/

WORKDIR /app
ENTRYPOINT ["./main"]