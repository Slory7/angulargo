FROM golang as builder
ENV GOPROXY https://goproxy.io
RUN mkdir /build 
WORKDIR /build

ADD ./src ./src
COPY go.mod .
RUN go mod download

WORKDIR /build/src/services/api

RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/src/services/api/main /app/
#COPY --from=builder /build/src/services/api/config.json /app/
WORKDIR /app
CMD ["./main"]