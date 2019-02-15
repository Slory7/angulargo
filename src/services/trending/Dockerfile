FROM golang as builder
RUN mkdir /build 
ADD ./src /build/src
ENV GOPATH $GOPATH:/build
WORKDIR /build/src
#RUN apk add git 
RUN go get -d ./...
WORKDIR /build/src/services/trending
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/src/services/trending/main /app/
COPY --from=builder /build/src/services/trending/config.json /app/
WORKDIR /app
CMD ["./main"]