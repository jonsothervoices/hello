package stringexercises

import "fmt"

//FizzBuzz Write a program that prints the numbers from 1 to 100;
//for multiples of 3 print "Fizz";
//for multiples of 5 print "Buzz";
//for multiples of 3 and 5 print "FizzBuzz"
func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("FizzBuzz ")
		case i%3 == 0:
			fmt.Printf("Fizz ")
		case i%5 == 0:
			fmt.Printf("Buzz ")
		default:
			fmt.Printf("%v ", i)
		}
	}
}
