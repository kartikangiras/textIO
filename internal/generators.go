package internal

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

func generatePassword(input)
