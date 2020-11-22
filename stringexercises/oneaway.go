package stringexercises

import "math"

//three types of edits: insert, remove, or replace a character.
//Given two strings, writa a function to check if they are one edit away.
//oneAway is O(n)
func oneAway(a, b string) bool {
	//early exit opportunity
	//length difference greater than 2 requires more than 1 edit
	if math.Abs(float64(len(b)-len(a))) > 1 {
		return false
	}

	for i, char1 := range a {
		//check length
		if len(b) <= i {
			return b == a[:i]
		}

		//found an edit
		if char1 != rune(b[i]) {
			//are we at the end of either a or b?
			if len(b) <= i+1 || len(a) <= i+1 {
				//only one edit, on last char
				return len(b) == len(a)
			}
			//insert case
			if a[i:] == b[i+1:] {
				return true
			}
			//remove case
			if a[i+1:] == b[i:] {
				return true
			}
			//replace case
			if a[i+1:] == b[i+1:] {
				return true
			}
			return false
		}
	}
	return true
}
