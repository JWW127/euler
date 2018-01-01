package main

import "fmt"

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?


//------------------------------solution 1 using J.F Sebastian & Euclids Algo.
func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}


func gcd(a, b int64) int64 {
	c := a % b

	if c == 0 {
		return b;		
	}

	return gcd(b, c)
}


func main() {
	var a, i int64 = 1, 2

	for ; i <= 20; i++ {
		a = lcm(a, i)
	}
	fmt.Println(a)
}

////-------------------------------------------------------end solution 1

//----------------------------------------------Solution 2 will use prime factorization ????

