package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func marshalInterface(input string) ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to reterive data", err)
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
		return nil, fmt.Errorf("failed to retrieve data", err)
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
