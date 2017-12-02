package main

import "fmt"

// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143


// Trial devision algo design

//will simplify this later to include bool in solution

func main() {
		var n int64 = 600851475143
		var d int64 = 2
	
		for d <= n {
			if n%d == 0 {
				n = n / d 
			} else {
				if d > 2 {
					d += 2 
				} else {
					d = 3
				}
			}
		}
		fmt.Println(d)
}
