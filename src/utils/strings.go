package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// Generate random string URL encoded
func GenerateRandomUrlEncodedString(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer)[:length], nil
}
