package main

import "fmt"

func fib(n int) int {
  if n < 0 {
    return 0
  }

  return 4 * fib(n-3) + fib(n-6)

}

func main() {
  //sum := 0
  //a := 1
  //b := 1
  //c := a + b


  //for c < 4000000 {
    //sum = sum + c
    //a = b + c
    //b = c + a
    //c = a + b
  //}

  sumfunc := fib(4000000)

  //fmt.Printf("%d\n", sum)
  fmt.Printf("%d\n", sumfunc)
}
