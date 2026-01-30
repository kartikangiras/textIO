package internal

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"

	"github.com/google/uuid"
)

func GenerateSHA256(input string) (string, error) {
	hash := sha256.Sum256([]byte(input))
	hashstring := hex.EncodeToString(hash[:])

	return hashstring, nil
}

func GenerateUUID() (string, error) {
	id := uuid.New()

	return (id.String()), nil
}

func GeneratePassword(length int) (string, error) {
	const charset = "0987654321qwertyuioplkjhgfdsazxvbnmx/;-+=!@#$%^&*"
	password := make([]byte, length)

	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]

	}
	return string(password), nil
}
