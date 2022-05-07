package rng

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"log"
	mathRand "math/rand"
	"temppaste/pkg/errorskit"
)

// InitRNG is a wrapper for SeedRNG. It exits the application on error.
func InitRNG() {
	err := SeedRNG()
	if err != nil {
		log.Fatalln(err)
	}
}

// SeedRNG seeds the math/rand RNG according to: https://stackoverflow.com/a/54491783.
// It should be initialized as soon as possible in the application.
func SeedRNG() error {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		return errorskit.Wrap(err, "cannot seed math/rand")
	}
	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	return nil
}
