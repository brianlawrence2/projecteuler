package main

import (
	"fmt"
)

func divisable(n int, a int) bool {

	for i := 2; i <= a; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func smallestDivisor(n int) int {

	i := 1

	for {
		if divisable(i, n) {
			return i
		}
		i++
	}
}

func main() {

	sDivisor := 0

	sDivisor = smallestDivisor(20)

	fmt.Printf("%d", sDivisor)
}
