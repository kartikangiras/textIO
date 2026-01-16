package internal

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func GetTextStats(input string) (int, int, int, int, error) {
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

func CleanUpText(input string) (string, error) {
	cleaned := strings.TrimSpace(input)

	collapsed := regexp.MustCompile(`\s+`).ReplaceAllString(cleaned, " ")

	return collapsed, nil
}

func ConvertCase(text, switchCase string) string {

	switch switchCase {

	case "uppercase":
		return strings.ToUpper(text)

	case "lowercase":
		return strings.ToLower(text)

	case "sentence":
		if text == "" {
			return ""
		}
		runes := []rune(text)
		runes[0] = unicode.ToUpper(runes[0])
		rest := strings.ToLower(string(runes[1:]))
		return string(runes[0]) + rest

	case "title":
		words := strings.Fields(text)
		for i, word := range words {
			if len(word) > 0 {
				runes := []rune(word)
				runes[0] = unicode.ToUpper(runes[0])
				words[i] = string(runes)
			}
		}
		return strings.Join(words, "")

	case "camelcase":
		text := strings.ToLower(text)
		reg := regexp.MustCompile("[-_ ]([a-z])")
		return reg.ReplaceAllStringFunc(text, func(match string) string {
			return strings.ToUpper(string(match[1]))
		})

	case "pascalcase":
		camel := ConvertCase(text, "camelcase")

		if camel == "" {
			return ""
		}
		runes := []rune(camel)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)

	case "snakecase":
		text = strings.ToLower(text)
		reg := regexp.MustCompile("[- ]")
		text = reg.ReplaceAllString(text, "_")
		reg = regexp.MustCompile("[^a-z0-9_]")
		return reg.ReplaceAllString(text, "")

	case "kebabcase":
		text = strings.ToLower(text)
		reg := regexp.MustCompile("[_ ]")
		text = reg.ReplaceAllString(text, "-")
		reg = regexp.MustCompile("[^a-z0-9_]")
		return reg.ReplaceAllString(text, "")

	case "constant-case":
		return strings.ToUpper(ConvertCase(text, "snakecase"))

	default:
		return text
	}
}
