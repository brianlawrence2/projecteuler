package main

import "fmt"

func IsPal(n int) bool {
  reverse := 0
  temp := n

  for temp > 0 {
    reverse = reverse * 10
    reverse = reverse + temp % 10
    temp = temp / 10
  }

  if n == reverse {
    return true
  } else {
    return false
  }
}

func main() {
  pal := 0

  for i := 100; i < 1000; i++ {
    for j := i; j < 1000; j++ {
      palCan := i * j
      if IsPal(palCan) && palCan > pal {
        pal = palCan
      }
    }
  }

  fmt.Printf("%d", pal)
}
