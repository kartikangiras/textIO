package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func MarshalInterface(input string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return "", fmt.Errorf("invalid JSON: %v", err)
	}
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("formatting failed: %v", err)
	}
	return string(bytes), nil
}

func KvJson(input string) (string, error) {
	lines := strings.Split(input, "\n")
	data := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		index := strings.IndexAny(line, ":=")

		if index != -1 {
			rawKey := strings.TrimSpace(line[:index])
			rawValue := strings.TrimSpace(line[index+1:])

			key := strings.Trim(rawKey, `"'`)
			value := strings.Trim(rawValue, `"'`)

			data[key] = value
		}
	}

	if len(data) == 0 {
		return "", fmt.Errorf("no valid key=value pairs found")
	}

	kv, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(kv), nil
}

func MinifyCSS(input string) (string, error) {
	regex := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	minified := regex.ReplaceAllString(input, "")

	minified = strings.ReplaceAll(minified, "\n", "")
	minified = strings.ReplaceAll(minified, "\t", "")

	minified = regexp.MustCompile(`\s+`).ReplaceAllString(minified, " ")
	minified = regexp.MustCompile(`\s*([{}:;,])\s*`).ReplaceAllString(minified, "$1")

	return strings.TrimSpace(minified), nil
}

func Encodebase64(input string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded, nil
}

func Decodebase64(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("invalid base64 input")
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
