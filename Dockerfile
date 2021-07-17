FROM golang:1.15

LABEL version="1.0.0"

WORKDIR /go/src/github.com/ardafirdausr/kaseer
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/kaseer cmd/kaseer/*.go

ENTRYPOINT ["/go/bin/kaseer"]