package main

import (
	"embed"
	"temppaste/pkg/rng"
)

//go:embed public/*
var publicFolder embed.FS

func main() {
	rng.InitRNG()
	startFiber(publicFolder)
}
