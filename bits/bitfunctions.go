package bits

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getBit(n, b int) bool { return (1<<b)&n != 0 }

func setBit(n, b int) int { return (1 << b) | n }

func clearBit(n, b int) int { return n & (^(1 << b)) }

func updateBit(n, b int, v bool) int {
	if v {
		return setBit(n, b)
	}
	return clearBit(n, b)
}

//5.1 Insertion: Given two 32-bit numbers M and N, and bit position i, write a function to insert M into N such that M starts at bit i.
//4:17-5:07~~50 minutes
func insertBits(n, m, i int) int {
	//create bool slice
	mBin := strconv.FormatInt(int64(m), 2)
	for j, v := range mBin {
		if v == '1' {
			n = setBit(n, i-j)
		} else {
			n = clearBit(n, i-j)
		}
	}
	return n
}

//5.2 Binary to String: Given a real number from 0 to 1, print the binary representation. If the number cannot be represented accurately in binary with at most 32 characters, print ERROR.

//5.3 Flip bit to Win: Given an integer, find the length of the longest sequece of 1s you could create by flipping 1 bit.

func flipToWin(n int) int {
	ans := 0
	nInt := int64(n)
	nBin := strconv.FormatInt(nInt, 2)
	if !strings.Contains(nBin, "0") {
		return len(nBin)
	}
	for i, v := range nBin {
		if v == '0' {
			if zeroNextToOne(nBin, i) {
				sum := lSum(nBin, i) + rSum(nBin, i)
				fmt.Printf("lsum %v, rsum %v\n", lSum(nBin, i), rSum(nBin, i))
				if sum > ans {
					ans = sum
				}
			}
		}
	}
	return ans + 1
}

func zeroNextToOne(s string, i int) bool {
	if i == 0 {
		if len(s) > 1 {
			return s[i+1] == '1'
		}
		return true
	}
	if i == len(s)-1 {
		if len(s) > 1 {
			return s[i-1] == '1'
		}
		return true
	}
	return s[i+1] == '1' || s[i-1] == '1'
}

func lSum(s string, n int) int {
	ret := 0
	for i := n - 1; i >= 0 && s[i] != '0'; i-- {
		ret++
	}
	return ret
}

func rSum(s string, n int) int {
	ret := 0
	for i := n + 1; i < len(s) && s[i] != '0'; i++ {
		ret++
	}
	return ret
}

//5.4 Next Number: Given a positive integer, print the next smallest and the next largest number that have the same number of 1 bits in their binary representation.

//smallest number bigger--O(n) where n is the langth of bin rep of n
//12:53~~1:30
func nextLNum(n int64) (int64, error) {
	//check max
	if n >= math.MaxInt64-1 {
		return 0, fmt.Errorf("error: %v out of range", n)
	}
	//bin rep
	nBin := "0" + strconv.FormatInt(n, 2)
	//find last 01
	i := strings.LastIndex(nBin, "01")
	//no 01?
	if i == -1 {
		return 0, fmt.Errorf("error: %v has all zeros: %v", n, nBin)
	}
	//flip last 01
	nBin = fmt.Sprintf("%v10%v", nBin[:i], nBin[i+2:])
	return strconv.ParseInt(nBin, 2, 64)
}

//largest number smaller
func nextSNum(n int64) (int64, error) {
	//check min(2)
	if n < 2 {
		return 0, fmt.Errorf("error: %v out of range", n)
	}
	//bin rep
	nBin := strconv.FormatInt(n, 2)
	//find last 10
	i := strings.LastIndex(nBin, "10")
	//no 10?
	if i == -1 {
		return 0, fmt.Errorf("error: %v has no nextSNum: %v", n, nBin)
	}
	//flip last 10
	nBin = fmt.Sprintf("%v01%v", nBin[:i], nBin[i+2:])
	return strconv.ParseInt(nBin, 2, 64)
}

//5.5: Debugger:
//((n&(n-1))==0)
//n==50
//n-1==49
//110010 &
//110001 == 0 ?
//110010 &
//001101 == 0
//when is n-1 the compliment of n?
//64==1000000
//63==0111111
//((n&(n-1))==0) iff n%2==0 ie. n is a power of 2
//this is a power of 2 check

//5.6: Conversion: Write a function to determine the number of bits you would need to flip to convert integer A to integer B (assume positive A and B).

//2:13~~2:21
func conv(a, b int64) int {
	s := strconv.FormatInt(a^b, 2)
	return strings.Count(s, "1")
}

//5.7:Pairwise swap: Write a function to swap odd and even bits in an integer with as few instructions as possible
//2:28~~3:30
func pairSwap(a int64) int64 {
	//maskL:shift a left 1
	//maskR:shift a right 1
	s := strconv.FormatInt(int64(a), 2)
	fmt.Printf("s is %v\n", s)
	len := len(s) - 1
	maskL := fmt.Sprintf("%v0", s)
	maskR := fmt.Sprintf("0%v", s[:len])
	fmt.Printf("maskL is %v\n", maskL)
	fmt.Printf("maskR is %v\n", maskR)
	//make new slice s with 0s length binrep of a
	for i := range s {
		//range check
		if i == len {
			s = fmt.Sprintf("%v%v", s[:i], maskL[i])
			fmt.Printf("i==len: s is now %v\n", s)
			continue
		}
		//if even, assign s[i] to maskL[i]
		if i%2 == 0 {
			s = fmt.Sprintf("%v%v%v", s[:i], maskL[i], s[i+1:])
			fmt.Printf("i is even: s is now %v\n", s)
			continue
		}
		//if index is odd, assign new[i] to maskR[i]
		s = fmt.Sprintf("%v%v%v", s[:i], maskR[i], s[i+1:])
		fmt.Printf("i is odd: s is now %v\n", s)
	}
	ret, _ := strconv.ParseInt(s, 2, 64)
	fmt.Printf("returning %v\n", ret)
	return ret
}

//isolate even and odds with prdefined masks???
//why the hell are 4s and 9s showing up??
