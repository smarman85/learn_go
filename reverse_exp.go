package main

import (
  "fmt"
)

func reverse(input string) string {
  chars := []rune(input)
  for i, j := 0, len(chars) - 1; i < j; i, j = i + 1, j - 1 {
    fmt.Println("i = %v", i)
    fmt.Println("j = %v", j)
    chars[i], chars[j] = chars[j], chars[i]
    fmt.Println(string(chars))
    fmt.Println("******************")
  }
  return "end"
}

func main() {
  reverse("abcdefghijklmnopqrstuvwxyz")
}
