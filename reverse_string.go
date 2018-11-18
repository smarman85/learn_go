package main

import "fmt"

func reverse(input string) string {
  // string are immutable in go
  // must be converted to a mutable array of runes []rune
  // perf revese on the run then recase as a string
  // !! a rune is an alias for int32 a characheer (Unicode code point)
  chars := []rune(input)
  for i, j := 0, len(chars)-1; i < j; i, j = i + 1, j - 1 {
    chars[i], chars[j] = chars[j], chars[i]
  }
  return string(chars)
}

func main() {
  fmt.Printf("%v\n", reverse("abcdefg"))
}
