<h1 align="center">PoS</h1>

<p align="center">
  <img src="https://travis-ci.com/ardafirdausr/go-pos.svg?branch=main" alt="Build Status">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License">  
  <img src="https://img.shields.io/github/v/release/ardafirdausr/go-pos.svg?style=flat" alt="Release Version">
</p>

Point of Sales. This app is simple implementation is point of sales using GO, Echo framework, and MySQL. 

## Setup
1. Copy .env.example file to .env
2. Fill the environment configurations
3. Get the module dependencies by running `go get ./...`

### Run The App
`go run ./cmd/pos/main.go`

### Test The App
`go test -v ./...`

### Build The App
`go build -o ./bin/pos ./cmd/pos/main.go`