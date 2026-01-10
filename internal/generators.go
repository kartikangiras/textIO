package internal

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/google/uuid"
)

func generateSHA256(input string) ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data: %v", err)
	}

	hash := sha256.Sum256([]byte(input))
	hashstring := hex.EncodeToString(hash[:])

	return []byte(hashstring), nil
}

func generateUUID() ([]byte, error) {
	id := uuid.New()

	return []byte(id.String()), nil
}

func generatePassword(length int) (string, error) {
	const charset = "abcdermcpxrnMAWIUE4RB3O2QSWD0192837465473OQI2938475RYHDJSNMZKj92837h4rdyenusdhbfuyswniaconst"
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
