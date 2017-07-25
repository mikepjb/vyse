package main

import "fmt"

func prompt() {
  var result string;

  fmt.Printf("Vyse v0\n")

  for true {
    fmt.Printf(">>> ")
    fmt.Scanln(&result)
    fmt.Printf(result + "\n")
  }
}

func main() {
  prompt()
}
