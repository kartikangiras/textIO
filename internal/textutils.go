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

func ConvertCase(input string) (string, error) {
	words := splitIntoWords(input)

	switch strings.ToLower(caseType) {
	case "upper":
		return strings.ToUpper(strings.Join(words, " "))

	case "lower":

		return strings.ToLower(strings.Join(words, " "))

	case "pascal":

		var sb strings.Builder
		for _, w := range words {
			sb.WriteString(capitalize(w))
		}
		return sb.String()

	case "camel":

		var sb strings.Builder
		for i, w := range words {
			if i == 0 {
				sb.WriteString(strings.ToLower(w))
			} else {
				sb.WriteString(capitalize(w))
			}
		}
		return sb.String()

	default:
		return input
	}
}
