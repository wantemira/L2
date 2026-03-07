// Package main является точкой входа в приложение.
package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
)

func main() {
	str := "a4bc2d5e"
	result, err := UnpackString(str)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	} 
	fmt.Printf("My string: %s. Result of unpacking: %s",str, result)

}

func UnpackString(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result strings.Builder

	var escaped, hasPrev bool
	var prev rune

	for _, r := range s {
		if escaped {
			result.WriteRune(r)
			prev = r
			hasPrev = true
			escaped = false
			continue
		}
		if r == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if !hasPrev {
				return "", ErrInvalidString
			}

			count := int(r - '0')
			if count == 0 {
				str := result.String()
				result.Reset()
				result.WriteString(str[:len(str)-len(string(prev))])
				hasPrev = false
				continue
			}

			for i := 0; i < count-1; i++ {
				result.WriteRune(prev)
			}
			continue
		}
		result.WriteRune(r)
		prev = r
		hasPrev = true
	}

	if escaped {
		return "", ErrInvalidString
	}

	return result.String(), nil
}
