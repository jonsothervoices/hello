package stringexercises

//Write a method to replace all spaces in a string with"%20"
//
func urlify(s string) string {
	for i, v := range s {
		if v == ' ' {
			if i < len(s)-1 {
				return s[:i] + "%20" + urlify(s[i+1:])
			}
			return s[:i] + "%20"
		}
	}
	return s
}
