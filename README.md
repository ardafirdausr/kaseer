<h1 align="center">Kaseer</h1>

<p align="center">
  <img src="https://travis-ci.com/ardafirdausr/kaseer.svg?branch=main" alt="Build Status">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License">  
  <img src="https://img.shields.io/github/v/release/ardafirdausr/kaseer.svg?style=flat" alt="Release Version">
</p>

Kaseer is a simple Point of Sales implementation using GO, Echo framework, and MySQL. 

## Setup
1. Copy .env.example file to .env
2. Fill the environment configurations
3. Get the module dependencies by running `go get ./...`

### Run The App
`go run ./cmd/kaseer/main.go`

### Test The App
`go test -v ./...`

### Build The App
`go build -o ./bin/kaseer ./cmd/kaseer/main.go`