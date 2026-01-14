package internal

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func marshalInterface(input string) ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to reterive data: %v", err)
	}

	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		log.Fatalf("error unmarshaling json: %v", err)
	}
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("error indentation of json: %v", err)
	}
	return bytes, nil
}

func kvJson(input string) ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data: %v", err)
	}
	entries := strings.Split(input, "\n")[0]
	data := make(map[string]string)

	delimiters := ":,="
	index := strings.IndexAny(entries, delimiters)

	if index != -1 {
		keypart := entries[:index]
		valuepart := entries[index+1:]

		key := strings.TrimSpace(keypart)
		value := strings.TrimSpace(valuepart)

		data[key] = value
		kv, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			return nil, fmt.Errorf("error in key value pair: %v", err)
		}
		return kv, nil
	}

	return nil, fmt.Errorf("no valid key-value pair found")
}

func minifyCSS(input string) ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read the data: %v", err)
	}

	regex := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	minified := regex.ReplaceAllString(input, "")

	minified = strings.ReplaceAll(minified, "\n", "")
	minified = strings.ReplaceAll(minified, "\t", "")
	minified = regexp.MustCompile(`\s+`).ReplaceAllString(minified, " ")
	minified = strings.TrimSpace(minified)

	return []byte(minified), nil
}

func encodebase64(input string) (any, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return nil, fmt.Errorf("failed to retreive the data: %v", err)
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	return encoded, nil
}

func decodebase64(input string) (any, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data: %v", err)
	}
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %v", err)
	}
	return decoded, nil
}

func encodeurl(input string) (any, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data: %v", err)
	}

	encoder := url.QueryEscape(input)

	return encoder, nil
}

func decodeurl(input string) (any, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data: %v", err)
	}

	decoder, err := url.QueryUnescape(input)

	if err != nil {
		return nil, fmt.Errorf("failed to decode url")
	}

	return decoder, nil
}
