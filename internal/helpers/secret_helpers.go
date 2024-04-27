package helpers

import (
	"math/rand"
	"time"
)

type SecretHelpers struct {
}

func NewSecretHelpers() *SecretHelpers {
	return &SecretHelpers{}
}

func (s *SecretHelpers) GenerateSecret(length int) string {
	src := rand.NewSource(time.Now().UTC().UnixNano())

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-!@#$%^&*()+=?/>.<,"

	result := make([]byte, length)
	for i := range result {
		result[i] = letterBytes[src.Int63()%int64(len(letterBytes))]
	}
	return string(result)
}
