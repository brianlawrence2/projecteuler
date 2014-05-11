package main

import (
	"fmt"
	"math"
)

func main() {

	k := 20
	N := 1
	i := 0
	check := true
	limit := math.Sqrt(float64(k))
	p := []int{2, 3, 5, 7, 11, 13, 17, 19}
	for i < 8 {
		a := 1
		if check {
			if float64(p[i]) <= limit {
				a = int(math.Floor(math.Log(float64(k)) / math.Log(float64(p[i]))))
				fmt.Printf("p: %d a: %d\n", p[i], a)
			} else {
				check = false
			}
		}
		N = N * int(math.Pow(float64(p[i]), float64(a)))
		fmt.Printf("p: %d N: %d\n", p[i], N)
		i++
	}

	fmt.Printf("%d\n", N)
}
