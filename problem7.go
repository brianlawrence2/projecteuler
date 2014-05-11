package main

import (
	"fmt"
	"math"
)

func sqrt(n int) int {
	return int(math.Floor(math.Sqrt(float64(n))))
}

func isPrime(n int) bool {
	switch {
	case n == 2:
		return true
	case n%2 != 0:
		for i := 3; i <= sqrt(n); i += 2 {
			if n%i == 0 {
				return false
			}
		}
	default:
		return false
	}
	return true
}

func main() {

	p := 1
	count := 0
	limit := 10001

	for count <= limit {
		if isPrime(p) {
			count++
		}
		if count < limit {
			p += 2
		}
	}

	fmt.Printf("%dth prime: %d", limit, p)
}
