package main

import "strings"

func Parse(s string) []string {
	if s == "" {
		return []string{}
	}

	return strings.Fields(s)
}
