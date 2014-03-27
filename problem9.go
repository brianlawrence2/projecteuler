package main

import (
  "fmt"
  "math"
)

type Abc struct {
  a int
  b int
  c int
}

func main() {

  lim := 1000
  t := 0
  can := make([]*Abc, 0)
  p := 0

  for i := 3; i < lim; i++ {
    for j := 1; j < i; j++ {
      for k := 1; k < j; k++ {
        t = i + j + k
        if t == lim {
          abc := new(Abc)
          abc.a = k
          abc.b = j
          abc.c = i
          can = append(can, abc)
        }
      }
    }
  }

  for _, i := range can {
    if math.Pow(float64(i.a), float64(2)) + math.Pow(float64(i.b), float64(2)) == math.Pow(float64(i.c), float64(2)) {
      p = i.a * i.b * i.c
    }
  }

  fmt.Printf("%d", p)
}
