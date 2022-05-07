package main

import (
	"temppaste/pkg/rng"
)

func main() {
	rng.InitRNG()
	startFiber()
}
