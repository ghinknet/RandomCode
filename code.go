package RandomCode

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Options is the option struct to generate random code
type Options struct {
	Digit         int    // Digit of the random code
	UseNumbers    bool   // Use numbers or not
	UseLowercase  bool   // Use lowercase or not
	UseUppercase  bool   // Use uppercase or not
	UseSymbols    bool   // Use symbols or not
	CustomSymbols string // Charset of custom symbols
	ExcludeChars  string // Charset to exclude
}

// DefaultSymbols is the default charset of symbols
const DefaultSymbols = "!@#$%^&*?"

// Code generates a random code
func Code(config Options) string {
	if config.Digit <= 0 {
		config.Digit = 6
	}

	// Build char set
	charSet := buildCharSet(config)
	if len(charSet) == 0 {
		charSet = "0123456789" // Rollback to default number charset
	}

	b := make([]byte, config.Digit)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			// Rollback to backup method
			b[i] = charSet[i%len(charSet)]
			continue
		}
		b[i] = charSet[n.Int64()]
	}
	return string(b)
}

// buildCharSet builds a charset for code generation
func buildCharSet(config Options) string {
	var charSet strings.Builder

	if config.UseNumbers {
		charSet.WriteString("0123456789")
	}
	if config.UseLowercase {
		charSet.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	if config.UseUppercase {
		charSet.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if config.UseSymbols {
		if config.CustomSymbols != "" {
			charSet.WriteString(config.CustomSymbols)
		} else {
			charSet.WriteString(DefaultSymbols)
		}
	}

	// Exclude specific char
	result := charSet.String()
	if config.ExcludeChars != "" {
		result = removeChars(result, config.ExcludeChars)
	}

	return result
}

// removeChars Remove specific chars from charset
func removeChars(source, charsToRemove string) string {
	var result strings.Builder
	for _, char := range source {
		if !strings.ContainsRune(charsToRemove, char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// Number generates a random code with num only
func Number(digit int) string {
	return Code(Options{
		Digit:      digit,
		UseNumbers: true,
	})
}

// Alpha generates a random code with alpha only
func Alpha(digit int, useLower, useUpper bool) string {
	return Code(Options{
		Digit:        digit,
		UseLowercase: useLower,
		UseUppercase: useUpper,
	})
}

// Mixed generates a random code with mixed charset
func Mixed(digit int, useNumbers, useLower, useUpper, useSymbols bool) string {
	return Code(Options{
		Digit:        digit,
		UseNumbers:   useNumbers,
		UseLowercase: useLower,
		UseUppercase: useUpper,
		UseSymbols:   useSymbols,
	})
}
