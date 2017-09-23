package main

import (
	"strings"
)

func sanitizeString(input string) (output string) {
	output = strings.TrimSpace(input)
	output = strings.ToLower(output)
	output = strings.Trim(input, "\"")
	return output
}
