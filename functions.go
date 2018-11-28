package main

import (
  "fmt"
)

func SayHelloToMe(firstName, lastName string, age int) {
  fmt.Printf("Hello, %s %s!\n", firstName, lastName)
  fmt.Printf("You are %d\n", age)
}

func main() {
  SayHelloToMe("Sean", "M", 33)
}
