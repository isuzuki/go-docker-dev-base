FROM golang:1.13-alpine as build

RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/go-sql-driver/mysql

ADD . /go/src/server
RUN go install server

FROM alpine
COPY --from=build /go/bin/server .
EXPOSE 8080

CMD ["./server"]