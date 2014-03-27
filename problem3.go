package main

import (
    "fmt"
    "math"
)

func isPrime(n int) bool {
  switch {
    case n == 2: return true
    case n % 2 != 0:
      for i := 3; i < sqrt(n); i+=2 {
        if n % i == 0 {
          return false
        }
      }
    default: return false
  }
  return true
}

func inSlice(n int, f []int) bool {
  for i := range f {
    if n == f[i] {
      return true
    }
  }
  return false
}

func max(f []int) int {
  p := 0
  m:= 0

  for i := range f {
     if f[i] > p {
       m = f[i]
     }
  }
  return m
}

func sqrt(n int) int {
  return int(math.Sqrt(float64(n)))
}

func factor(n int, f []int) []int {

  for i := 2; i < sqrt(n); i++ {
    mod := n % i
    if mod == 0 {
      if isPrime(i) {
        if !inSlice(i, f) {
          f = append(f, i)
        }
      } else {
        f = append(f, factor(i, f)...)
      }
    }
  }
  return f
}


func main() {

  composite := 600851475143

  factors := make([]int, 1)

  factors = factor(composite, factors)

  m := max(factors)

  fmt.Printf("%s", factors)
  fmt.Printf("%d", m)
}
