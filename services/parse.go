package services

import "strings"

func ParsePath(s string) []string {
	if s == "" {
		return []string{}
	}

	return strings.Fields(s)
}

func ParseName(s string) string {
	runes := []rune(s)
	index := 0
	for i := len(runes) - 1; i > 0; i-- {
		if runes[i] == '/' {
			index = i + 1
			break
		}
	}

	return s[index:]
}
