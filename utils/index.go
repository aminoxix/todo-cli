package appUtils

import (
	"crypto/rand"
	"math/big"
)

func Utils() *big.Int {
	randN, _ := rand.Int(rand.Reader, big.NewInt(100000))
	return randN
}
