package main 

import "fmt"

func main () {
  return_val := ""
  for i := 1; i <= 100; i ++ {
    if i % 5 == 0 {
      return_val += "Bazz"
    }
    if i % 3 == 0 {
      return_val += "Fizz"
    }
    if i % 2 == 0 {
      return_val += "Buzz"
    }
    if return_val != "" {
      fmt.Println(return_val)
    }
    // rest return val to ""
    return_val = ""
  }
}
