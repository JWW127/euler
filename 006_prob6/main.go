package main

import "fmt"
import "math"


// The sum of the squares of the first ten natural numbers is,

// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,

// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.


func main() {
	var sumsquare = float64(1)
	var squaresum = float64(1)
	for i := float64(2); i <= 100; i++ {
			squaresum += i
			sumsquare += i * i
	}
	squaresum = math.Pow(squaresum, 2)
	fmt.Print(squaresum - sumsquare)
}