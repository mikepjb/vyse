package main

import (
  "fmt"
  "strings"
  "bufio"
  "io"
  "os"
)

func tokenize(code string) string {
  return strings.Replace(
    strings.Replace(code, "(", " ( ", -1), ")", " ) ", -1)
}

func prompt(in io.Reader, out io.Writer) error {
  fmt.Fprintf(out, "Vyse v0\n")
  fmt.Fprintf(out, ">>> ")

  scanner := bufio.NewScanner(in)
  for scanner.Scan() {
    result := scanner.Text()
    token := tokenize(result)

    fmt.Println(token)
    fmt.Fprintf(out, ">>> ")
  }

  return scanner.Err()
}

func main() {
  err := prompt(os.Stdin, os.Stdout)
  if err != nil {
    panic(err)
  }
}
