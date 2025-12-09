package services

import (
	"crypto/rand"
	"fmt"
)

var defaultAlphabet = []rune("0123456789abcdef")

func GeneratePublicID[T ~string](existingID T) (T, error) {
	if existingID != "" {
		return existingID, nil
	}

	id, err := newNanoID(8)
	if err != nil {
		return "", fmt.Errorf("error generating nano-id: %w", err)
	}

	return T(id), nil
}

func newNanoID(size int) (string, error) {
	if size < 1 || size > 30 {
		return "", fmt.Errorf("invalid size: %d", size)
	}

	bytes := make([]byte, size)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	id := make([]rune, size)

	alphabetLen := len(defaultAlphabet)

	for i := range size {
		id[i] = defaultAlphabet[int(bytes[i])%alphabetLen]
	}

	return string(id[:size]), nil
}
