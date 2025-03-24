package strings

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandom(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
