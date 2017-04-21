package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 3

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency {
	m := make(Frequency)
	a := regexp.MustCompile("[,:\n ]+")
	words := a.Split(phrase, -1)
	for _, word := range words {
		m[strings.Trim(strings.ToLower(word), "'.!@$%^&")]++
	}
	return m
}
