package main

import (
  "fmt"
)

func main() {
  var nums = [10]int{0,1,2,3,4,5,6,7,8,9}

  // print length of the array
  fmt.Println(len(nums))

  // print break line
  fmt.Println("**")

  // iterate over the array and print each index
  for i := 0; i < len(nums); i ++ {
    fmt.Println(nums[i])
  }

  // print break line
  fmt.Println("**")

  // print in revers, cause why not
  for i := len(nums) - 1; i >= 0; i -- {
    fmt.Println(nums[i])
  }
}
