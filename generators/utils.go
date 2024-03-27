package generators

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// contains checks if a string is present in a slice of strings.
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// singularize attempts to convert a plural noun to its singular form.
// This is a basic implementation and may not correctly handle all plural forms.
func singularize(word string) string {
	// Basic rules for regular plural forms
	if strings.HasSuffix(word, "ies") {
		return strings.TrimSuffix(word, "ies") + "y"
	} else if strings.HasSuffix(word, "ves") {
		return strings.TrimSuffix(word, "ves") + "f"
	} else if strings.HasSuffix(word, "s") {
		// Assumes words ending in 's' are plural
		return strings.TrimSuffix(word, "s")
	}

	// Return the word if no rules apply
	return word
}

// uppercaseFirst converts the first letter of the string to uppercase.
func uppercaseFirst(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
