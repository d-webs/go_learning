package main

import (
  "fmt"
)

type node struct {
  data int
  next *node
}

type list struct {
  head *node
  length int
}

// func (l list) append(n *node) {
//
// }
//
// func (l list) pop() *node {
//
// }


func main() {
  fmt.Println("Hello, World!")
}
