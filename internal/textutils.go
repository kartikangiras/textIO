package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func getTextStats(input string) (int, int, int, int, error) {
	characterCount := utf8.RuneCountInString(input)

	words := strings.Fields(input)
	wordCount := len(words)

	lineCount := strings.Count(input, "\n") + 1

	noSpaceCount := 0
	for _, r := range input {
		if r != ' ' && r != '\t' && r != '\n' {
			noSpaceCount++
		}
	}
	return characterCount, wordCount, lineCount, noSpaceCount, nil
}

func cleanUpText(input string) (string, error) {
	cleaned := strings.TrimSpace(input)

	collapsed := regexp.MustCompile(`\s+`).ReplaceAllString(cleaned, " ")

	return collapsed, nil
}

func convertCase(text, switchCase string) string {

	switch switchCase {
		
	case "uppercase":
		return strings.ToUpper(text)

	case "lowercase" : 
		return strings.ToLower(input)

	case "sentence" :
		if text == ""{
			return ""
		}
		runes := []rune(text)
		runes[0] := unicode.ToUpper(runes[0])
		rest := strings.ToLower(string(runes[1:]))
		return string(runes[0]) + rest

	case "title": 
		words := strings.Fields(text)
		for i, word := range words {
			if len(word) > 0 {
				runes := []rune(word)
				runes[0] := unicode.ToUpper(runes[0])
				words[i] := string(runes)
			}
		}
		return strings.Join(words, "")

	case ""
		
}
