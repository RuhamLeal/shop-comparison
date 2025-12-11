package services

import (
	"crypto/rand"
	"fmt"
	"strings"
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

func FormatCentsToBRL(cents int64) string {
	var brl int64
	if cents < 0 {
		brl = -cents / 100
	} else {
		brl = cents / 100
	}

	s := fmt.Sprintf("%d", brl)

	var b strings.Builder
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if count == 3 {
			b.WriteByte('.')
			count = 0
		}
		b.WriteByte(s[i])
		count++
	}

	bytes := []byte(b.String())
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return fmt.Sprintf("R$ %s,%02d", string(bytes), cents%100)
}
