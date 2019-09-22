package main

import (
	"math/rand"
	"time"
)

//GOARCH=wasm GOOS=js go build -o lib.wasm main.go
//go test -cpuprofile profile.out
//go tool pprof --web profile.out
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	select{}
}

