//Given a string, write a function to check if it is a permutation of a palindrome.
//Basically, can the string be rearranged to form a palendrome?

package stringexercises

import "strings"

//PalPermute has time of O(n)
//PalPermute is a palindrome permutation checker
func PalPermute(s string) bool {
	m := make(map[string]int)

	for _, v := range s {
		cleanv := strings.TrimSpace(strings.ToLower(string(v)))
		if len(cleanv) > 0 {
			m[cleanv]++
		}
	}

	foundodd := false
	for _, v := range m {
		if v%2 == 0 {
			continue
		}
		if foundodd {
			return false
		}
		foundodd = true
	}

	return true
}
