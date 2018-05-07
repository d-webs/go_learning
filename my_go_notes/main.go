package main

import "fmt"

func modifyInteger(arg *int) {
  // NB: "*" in params tells the function to create a pointer to its memory location. Without it we cannot mutate what was given

  // intVal := *ptrToInt

  intVal := *arg
  intVal = intVal * 2
  *arg = intVal

  // one linter: *arg *= 2
}

func main() {
  // breakfast := [...]string{"eggs", "bacon", "grits", "biscuits"}
  // ... tells copmiler to figure out how long the array is

  // breakfast[2] = "sausage"

  val := 123
  modifyInteger(&val)
  // NB: "&" creates a pointer to the variable

  fmt.Println(val)

  // fmt.Printf("%v\n", breakfast)
}

/*

NOTE:

Objects passed into functions are not mutable in Go


*/
