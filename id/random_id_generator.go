package id

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func GenerateRandomID() (string, error) {
	// Create a byte slice to store the random bytes
	idBytes := make([]byte, 16)

	// Read random bytes from the crypto/rand package
	_, err := io.ReadFull(rand.Reader, idBytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a hexadecimal string
	id := hex.EncodeToString(idBytes)

	return id, nil
}
