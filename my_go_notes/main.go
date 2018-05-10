package main

import "fmt"

func modifyInteger(arg *int) {
  // NOTE: "*" in params tells the function to create a pointer to its memory location. Without it we cannot mutate what was given

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

In order to pass in an object by reference, you must:
  pass it in like this: methodName(&object)
  have the method accept the object func growOneYear(c *cat) {}


In go, you can have methods on things that are not structs

You cannot define methods on built in Go types - you can make an alias for any type, however
  type nedInt int
    --> this makes an integer-like object (basically an alias)

array syntax
  myNums := [2]int{}

Casting:
  if you have a stack of objects that use the anything interface:
    var a2 anything
    var num2 int

    a2 = stack.pop()
    num2 = a2.(int) <-- castings
*/
