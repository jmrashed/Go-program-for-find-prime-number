package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
		fmt.Printf("%d\n", i)
	}
	return true
}
func main() {
	var input int // input integer
	fmt.Println("Enter input")
	fmt.Scanln(&input)

	if isPrime(input) {
		fmt.Printf("%d is a prime number \n", input)
	} else {
		fmt.Printf("%d is not a prime number \n", input)
	}
}
