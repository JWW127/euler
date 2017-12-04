package main

import (
	"fmt"
	"strconv"
)

//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
//Find the largest palindrome made from the product of two 3-digit numbers.
//
//	This one is hard! Basically looked at a bunch of C examples of this problem. This is a pretty good way to do it, but feels messy because im terrible. Will have to make it pretty later.


func isPal(n int64) bool {
	if n < 0 {
		n = -n
	}


	//converting to string does not seem right, but my brain wasn't working good enough for other solution
	s := strconv.FormatInt(n, 10)
	bound := (len(s) / 2) + 1
	for i := 0; i < bound; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	Largest := 0
	numI := 0
	numJ := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if Largest > product {
				break
			}
			if isPal(int64(product)) {
				Largest = product
				numI = i
				numJ = j
				continue
			}
		}
	}
	fmt.Println("answer: ",numI, "x",numJ,"=", Largest)
}