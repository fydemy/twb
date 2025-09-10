package helper

import (
	"crypto/rand"
	"encoding/base64"
)

func NewBase64URL() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
