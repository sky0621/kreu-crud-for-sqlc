FROM golang:1.20 as build

WORKDIR /go/src/app

COPY . .
RUN CGO_ENABLED=1 go build -o /go/bin/app
CMD ["/go/bin/app"]