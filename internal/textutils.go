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

	lineCount := 0
	if input != "" {
		lineCount = strings.Count(input, "\n") + 1
	}

	noSpaceCount := 0
	for _, r := range input {
		if !unicode.IsSpace(r) {
			noSpaceCount++
		}
	}
	return characterCount, wordCount, lineCount, noSpaceCount, nil
}

func CleanUpText(input string, action string) (string, error) {
	switch action {
	case "removeExtraSpaces":
		cleaned := strings.TrimSpace(input)
		collapsed := regexp.MustCompile(`\s+`).ReplaceAllString(cleaned, " ")
		return collapsed, nil

	case "removeLineBreaks":
		result := regexp.MustCompile(`\r?\n`).ReplaceAllString(input, " ")
		result = regexp.MustCompile(`\s+`).ReplaceAllString(result, " ")
		return strings.TrimSpace(result), nil

	case "removeAllSpaces":
		result := regexp.MustCompile(`\s`).ReplaceAllString(input, "")
		return result, nil

	case "trimLines":
		lines := strings.Split(input, "\n")
		var cleanedLines []string
		for _, line := range lines {
			cleanedLines = append(cleanedLines, strings.TrimSpace(line))
		}
		return strings.Join(cleanedLines, "\n"), nil

	default:
		return strings.TrimSpace(input), nil
	}
}

func ConvertCase(text, switchCase string) (string, error) {
	if text == "" {
		return "", nil
	}

	switch switchCase {

	case "uppercase":
		return strings.ToUpper(text), nil

	case "lowercase":
		return strings.ToLower(text), nil

	case "sentence":
		runes := []rune(strings.ToLower(text))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		return string(runes), nil

	case "title":
		words := strings.Fields(text)
		for i, word := range words {
			runes := []rune(strings.ToLower(word))
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
			}
			words[i] = string(runes)
		}
		return strings.Join(words, " "), nil

	case "camelcase":
		reg := regexp.MustCompile(`[-_]+`)
		text = reg.ReplaceAllString(text, " ")

		words := strings.Fields(text)
		for i, word := range words {
			if i == 0 {
				words[i] = strings.ToLower(word)
			} else {
				runes := []rune(strings.ToLower(word))
				if len(runes) > 0 {
					runes[0] = unicode.ToUpper(runes[0])
				}
				words[i] = string(runes)
			}
		}
		return strings.Join(words, ""), nil

	case "pascalcase":
		camel, _ := ConvertCase(text, "camelcase")
		runes := []rune(camel)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		return string(runes), nil

	case "snakecase":
		return toDelimited(text, '_'), nil

	case "kebabcase":
		return toDelimited(text, '-'), nil

	case "constant-case":
		snake := toDelimited(text, '_')
		return strings.ToUpper(snake), nil

	default:
		return text, nil
	}
}

func toDelimited(text string, delimiter rune) string {
	text = strings.TrimSpace(text)
	var result strings.Builder

	for i, r := range text {
		if i > 0 && unicode.IsUpper(r) {
			prev := rune(text[i-1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) {
				result.WriteRune(delimiter)
			}
		}
		result.WriteRune(r)
	}

	s := result.String()
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = reg.ReplaceAllString(s, string(delimiter))

	s = strings.Trim(s, string(delimiter))

	return strings.ToLower(s)
}
