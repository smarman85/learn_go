package main

import (
  "fmt"
  "strings"
)

func reverse_words(input string) string {
  // convert the string into an array of strings where each entry
  // is a word
  words := strings.Fields(input)
  // reverse loop over the array
  for i, j := 0, len(words) - 1; i < j; i, j = i + 1, j - 1 {
    words[i], words[j] = words[j], words[i]
  }
  // rebuild the string joinging with a space
  return strings.Join(words, " ")
}

func main() {
  fmt.Println(reverse_words("one two three"))
}
