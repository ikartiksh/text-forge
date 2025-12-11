package utils

import (
	"regexp"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ToUpperCase converts text to uppercase
func ToUpperCase(text string) string {
	return strings.ToUpper(text)
}

// ToLowerCase converts text to lowercase
func ToLowerCase(text string) string {
	return strings.ToLower(text)
}

// ToTitleCase converts text to title case
func ToTitleCase(text string) string {
	caser := cases.Title(language.English)
	return caser.String(text)
}

// ReverseText reverses the input text
func ReverseText(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// WordCount returns word, character, line, and paragraph counts
func WordCount(text string) map[string]int {
	text = strings.TrimSpace(text)

	// Count lines
	lines := strings.Split(text, "\n")
	lineCount := len(lines)

	// Count paragraphs (separated by blank lines)
	paragraphs := strings.Split(text, "\n\n")
	paragraphCount := 0
	for _, p := range paragraphs {
		if strings.TrimSpace(p) != "" {
			paragraphCount++
		}
	}

	// Count characters (with and without spaces)
	charCount := len(text)
	charNoSpaces := len(strings.ReplaceAll(strings.ReplaceAll(text, " ", ""), "\n", ""))

	// Count words
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

// TrimText trims whitespace from text
func TrimText(text string) string {
	return strings.TrimSpace(text)
}

// FindReplace finds and replaces text
func FindReplace(text, find, replace string, caseSensitive bool) string {
	if !caseSensitive {
		re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(find))
		return re.ReplaceAllString(text, replace)
	}
	return strings.ReplaceAll(text, find, replace)
}

// RemoveDuplicateLines removes duplicate lines from text
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

// SortLines sorts lines alphabetically
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

// ConvertCase converts between different case formats
func ConvertCase(text, caseType string) string {
	switch caseType {
	case "camelCase":
		return toCamelCase(text)
	case "PascalCase":
		return toPascalCase(text)
	case "snake_case":
		return toSnakeCase(text)
	case "kebab-case":
		return toKebabCase(text)
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
			// Check for camelCase transition
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
