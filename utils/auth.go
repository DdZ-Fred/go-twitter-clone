package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"strings"
)

type CodeVerifier struct {
	Value string
}

func (cv *CodeVerifier) CodeChallengePlain() string {
	return cv.Value
}

func (cv *CodeVerifier) CodeChallengeS256() string {
	h := sha256.New()
	h.Write([]byte(cv.Value))
	return base64Encode(h.Sum(nil))
}

func base64Encode(bytes []byte) string {
	encoded := base64.StdEncoding.EncodeToString(bytes)
	encoded = strings.Replace(encoded, "+", "-", -1)
	encoded = strings.Replace(encoded, "/", "_", -1)
	encoded = strings.Replace(encoded, "=", "", -1)
	return encoded
}

func GenerateCodeVerifier() *CodeVerifier {
	// Initialize random slice of bytes
	randomBytes := make([]byte, 40, 40)
	for i := 0; i < 40; i++ {
		randomBytes[i] = byte(rand.Intn(255))
	}

	return &CodeVerifier{
		Value: base64Encode(randomBytes),
	}
}
