package main

import "fmt"

func findPrimes(lim int) []int {
	can := make([]int, lim+1)

	for i := 0; i <= lim; i++ {
		can[i] = 1
	}

	can[0] = 0
	can[1] = 0

	for i := 2; i <= lim; i++ {
		if can[i] == 1 {
			j := i + i
			for j <= lim {
				can[j] = 0
				j = j + i
			}
		}
	}

	return can
}

func main() {
	can := findPrimes(2000000)

	sum := 0
	for p, b := range can {
		if b == 1 {
			sum = sum + p
		}
	}

	fmt.Printf("%d\n", sum)
}
