package main

import (
	"fmt"
)

// isPrime checks if a given number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}

	// Print all prime numbers in a range
	fmt.Println("Prime numbers between 1 and 50:")
	for i := 1; i <= 50; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
