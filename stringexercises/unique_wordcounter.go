package stringexercises

import "strings"

//O(n)
//discard empty strings
func uniqueWordCounter(words []string) int {
	m := make(map[string]bool)

	for _, v := range words {
		if v == "" {
			continue
		}
		m[strings.ToLower(v)] = true
	}
	return len(m)
}
