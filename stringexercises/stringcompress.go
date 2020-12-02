package stringexercises

import "fmt"

//Implement a basic string "compression" using the counts of repeated characters, such that the output read reads "Letter1Count1Letter2Count2...". If the compressed string would not become smaller, return the original string.

func stringCompress(s string) string {
	counter := 1
	c := ""
	newString := ""
	for i, r := range s {
		c = string(r)
		if i == 0 {
			continue
		}
		if s[i] != s[i-1] {
			newString += (string(s[i-1]) + fmt.Sprintf("%v", counter))
			counter = 1
			continue
		}
		counter++
	}
	newString = (newString + fmt.Sprintf("%v%v", c, counter))
	if len(s) <= len(newString) {
		return s
	}
	return newString
}
