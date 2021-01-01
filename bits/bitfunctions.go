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
//11:40~~12:17
func pairSwap(a int64) int64 {
	maskE, _ := strconv.ParseInt("101010101010101010101010101010101010101010101010101010101010101", 2, 64)
	maskO, _ := strconv.ParseInt("010101010101010101010101010101010101010101010101010101010101010", 2, 64)
	evens := (maskE & a) << 1
	odds := (maskO & a) >> 1
	fmt.Printf("evens is %v\nodds is %v\n", evens, odds)
	return evens | odds
}

//5.8: Draw Line: a monochrome screen is stored as an array of bits, allowing eight consecutive pixels to be stored in one byte. The screen has width w (w%8==0). Write a function to draw a horizontal line from (x1,y) to (x2,y).

//12:57~~2:30
func drawLine(x1, x2, y, w int, screen []byte) {
	bitStart := y*w + x1
	bitEnd := y*w + x2
	for i := bitStart; i <= bitEnd; i++ {
		screen[i/8] = (1 << (7 - i%8)) | screen[i/8]
	}
}
