FROM golang:1.15

LABEL version="1.0.0"

WORKDIR /go/src/github.com/ardafirdausr/pos
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/pos cmd/pos/*.go

ENTRYPOINT ["/go/bin/pos"]