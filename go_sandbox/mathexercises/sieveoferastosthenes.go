package mathexercises

import (
	"math"
)

func SieveOfEratosthenes(max int) []int {
	s := make([]bool, max)
	for i := range s {
		s[i] = true
	}
	s[0] = false
	s[1] = false
	//iterate through all numbers from 2 to max (i)...
	for i := 2; i <= int(math.Sqrt(float64(max))); i++ {
		//...as long as s[i] is true
		if !s[i] {
			continue
		}
		//each iteration, we find all j, multiples of i, less than max...
		for j := i + 1; j < max; j++ {
			//...and set them to false
			if j%i == 0 {
				s[j] = false
			}
		}
	}
	//build and return int slice where indices are true
	iSlice := []int{}
	for i, v := range s {
		if v {
			iSlice = append(iSlice, i)
		}
	}
	return iSlice
}
