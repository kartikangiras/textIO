package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
)

func MarshalInterface(input string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		log.Fatalf("error unmarshaling json: %v", err)
	}
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("error indentation of json: %v", err)
	}
	return string(bytes), nil
}

func KvJson(input string) (string, error) {
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
			return "", fmt.Errorf("error in key value pair: %v", err)
		}
		return string(kv), nil
	}
	return "", fmt.Errorf("no valid key-value delimiter found")
}

func MinifyCSS(input string) (string, error) {
	regex := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	minified := regex.ReplaceAllString(input, "")

	minified = strings.ReplaceAll(minified, "\n", "")
	minified = strings.ReplaceAll(minified, "\t", "")
	minified = regexp.MustCompile(`\s+`).ReplaceAllString(minified, " ")
	minified = strings.TrimSpace(minified)

	return strings.TrimSpace(minified), nil
}

func Encodebase64(input string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	return encoded, nil
}

func Decodebase64(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("failed to decode:")
	}
	return string(decoded), nil
}

func Encodeurl(input string) (string, error) {
	encoder := url.QueryEscape(input)
	return encoder, nil
}

func Decodeurl(input string) (string, error) {
	decoder, err := url.QueryUnescape(input)
	if err != nil {
		return "", fmt.Errorf("failed to decode url")
	}
	return decoder, nil
}
