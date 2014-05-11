package main

import (
	"fmt"
)

func main() {

	s := 1000
	t := 0

	for a := 3; a <= (s-3)/3; a++ {
		for b := (a + 1); b <= (s-1-a)/2; b++ {
			c := s - a - b
			if c*c == a*a+b*b {
				t = a * b * c
			}
		}
	}

	fmt.Printf("%d", t)

}
