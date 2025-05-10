package config

import (
	"crypto/rand"
	"math/big"
)

func GenerateId() (uint, error) {
	// Generate random number between 10000 and 99999
	n, err := rand.Int(rand.Reader, big.NewInt(90000))
	if err != nil {
		return 0, err
	}
	return uint(n.Int64() + 10000), nil
}
