package main

import (
  "fmt"
)


func main() {
  firstList := linkedList{}
  
  fmt.Printf("Right now, our list looks like %T\n", firstList)

  firstList.insert(1)
  firstList.insert(2)
  firstList.insert(3)
  firstList.insert(4)
  firstList.insert(5)

  fmt.Printf("The head node has a value %v\n", firstList.head.data)

  current := firstList.head

  for current.next != nil {
    current = current.next
  }

  fmt.Printf("The tail of the linked list is %v\n", current.data)

  fmt.Printf("The head node has a value %v\n", firstList.head.data)
  fmt.Printf("Right now our list has a length of %v\n", firstList.length)
  fmt.Printf("When I pop the list, the value we get is %v\n", firstList.pop())

  myFirstQueue := queue{}
  myFirstQueue.push(5)
  myFirstQueue.push(6)
  
  newList := linkedList{}

  newList.insert(5)
  
  myQueue := newQueue()
  myQueue.push(1)
  myQueue.push(2)
  myQueue.push(3)
  fmt.Printf("The last value is %v\n", myQueue.pop())
}
