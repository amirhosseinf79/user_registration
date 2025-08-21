package pkg

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateNumericOTP(nDigits int) (string, error) {
	if nDigits <= 0 || nDigits > 9 {
		return "", fmt.Errorf("invalid digits")
	}
	mod := int64(1)
	for range nDigits {
		mod *= 10
	}
	nBig, err := rand.Int(rand.Reader, big.NewInt(mod))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0*d", nDigits, nBig.Int64()), nil
}
