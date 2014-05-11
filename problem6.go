package main

import (
	"fmt"
	"math"
)

func squareSum(n int) int {

	s := 0
	sq := 0

	for i := 0; i <= n; i++ {
		s += i
	}

	sq = int(math.Pow(float64(s), float64(2)))

	return int(sq)
}

func sumSquares(n int) int {

	s := 0

	for i := 0; i <= n; i++ {
		s += int(math.Pow(float64(i), float64(2)))
	}

	return s
}

func main() {

	d := 0
	s1 := squareSum(100000000)
	s2 := sumSquares(100000000)

	d = s1 - s2

	fmt.Printf("Dif: %d", d)
}
