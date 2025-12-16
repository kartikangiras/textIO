package utils

import (
	"regexp"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToUpperCase(text string) string {
	return strings.ToUpper(text)
}

func ToLowerCase(text string) string {
	return strings.ToLower(text)
}

func ToTitleCase(text string) string {
	caser := cases.Title(language.English)
	return caser.String(text)
}

func ReverseText(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func WordCount(text string) map[string]int {
	text = strings.TrimSpace(text)

	lines := strings.Split(text, "\n")
	lineCount := len(lines)

	paragraphs := strings.Split(text, "\n\n")
	paragraphCount := 0
	for _, p := range paragraphs {
		if strings.TrimSpace(p) != "" {
			paragraphCount++
		}
	}

	charCount := len(text)
	charNoSpaces := len(strings.ReplaceAll(strings.ReplaceAll(text, " ", ""), "\n", ""))

	words := strings.Fields(text)
	wordCount := len(words)

	return map[string]int{
		"words":              wordCount,
		"characters":         charCount,
		"charactersNoSpaces": charNoSpaces,
		"lines":              lineCount,
		"paragraphs":         paragraphCount,
	}
}

func TrimText(text string) string {
	return strings.TrimSpace(text)
}

func FindReplace(text, find, replace string, caseSensitive bool) string {
	if !caseSensitive {
		re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(find))
		return re.ReplaceAllString(text, replace)
	}
	return strings.ReplaceAll(text, find, replace)
}

func RemoveDuplicateLines(text string) string {
	lines := strings.Split(text, "\n")
	seen := make(map[string]bool)
	var result []string

	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

func SortLines(text string, ascending bool) string {
	lines := strings.Split(text, "\n")

	sort.Slice(lines, func(i, j int) bool {
		if ascending {
			return strings.ToLower(lines[i]) < strings.ToLower(lines[j])
		}
		return strings.ToLower(lines[i]) > strings.ToLower(lines[j])
	})

	return strings.Join(lines, "\n")
}

func ConvertCase(text, caseType string) string {
	switch caseType {
	case "camelCase":
		return toCamelCase(text)
	case "PascalCase":
		return toPascalCase(text)
	case "snake_case":
		return toSnakeCase(text)
	case "CONSTANT_CASE":
		return toConstantCase(text)
	default:
		return text
	}
}

func toCamelCase(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}
	result := strings.ToLower(words[0])
	for _, word := range words[1:] {
		result += strings.Title(strings.ToLower(word))
	}
	return result
}

func toPascalCase(s string) string {
	words := splitWords(s)
	var result string
	for _, word := range words {
		result += strings.Title(strings.ToLower(word))
	}
	return result
}

func toSnakeCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "_")
}

func toKebabCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "-")
}

func toConstantCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}
	return strings.Join(words, "_")
}

func splitWords(s string) []string {
	var words []string
	var currentWord strings.Builder

	for i, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if i > 0 && unicode.IsUpper(r) && unicode.IsLower(rune(s[i-1])) {
				if currentWord.Len() > 0 {
					words = append(words, currentWord.String())
					currentWord.Reset()
				}
			}
			currentWord.WriteRune(r)
		} else {
			if currentWord.Len() > 0 {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
		}
	}

	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words
}
