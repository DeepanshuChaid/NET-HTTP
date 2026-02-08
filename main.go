package main

import (
  "fmt"
  'github.com/DeepanshuChaid/NET-HTTP/internal/config"
)

func main () {
  config := config.MustLoad()
  
  fmt.Println("Hello World")
}