package main

import (
	"strings"
	"testing"
)

func TestRename(t *testing.T) {
	walkDir("hhe", []string{"IMG", "CSS", "JS"}, func(s string) string {
		return strings.ToLower(s)
	})
}
