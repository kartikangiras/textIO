package internal

import (
	"regexp"
	"strings"
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
