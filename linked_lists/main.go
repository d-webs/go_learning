package main

import (
  "fmt"
)

// STRUCT DEFINITIONS //////

type node struct {
  data int
  next *node
}

type linkedList struct {
  head *node 
  length int 
}

type queue struct {
  inStack *linkedList
  outStack *linkedList
}

// QUEUE METHODS /////

func (q *queue) push(val int) {
  if q.inStack == nil {
    q.inStack = &linkedList{}
  }

  q.inStack.insert(val)
}

func newQueue() queue {
  q := queue{}
  q.inStack, q.outStack = &linkedList{}, &linkedList{}

  return q
}

func (q *queue) pop() int {
  if q.outStack == nil {
    q.outStack = &linkedList{}
  }

  if q.outStack.isEmpty() {
    q.flip()
    return q.outStack.pop()
  } else {
    return q.outStack.pop()
  }
}

func (q *queue) flip() {
  for !q.inStack.isEmpty() {
    q.outStack.insert(q.inStack.pop())
  }
}

// func (q *queue) peek() int {
//   var topVal int

//   if q.outStack == nil && q.inStack == nil {
//     panic("The queue is empty")
//   } else if q.outStack.isEmpty() && q.inStack.isEmpty() {
//     panic("The queue is empty")
//   } else if q.outStack.isEmpty() {
//     current := q.inStack.head 

//     for current.next != nil {
//       current = current.next
//     }

//     topVal = current.data
//   } else {
//     topVal = q.outStack.head.data
//   }

//   return topVal
// }

// LINKED LIST METHODS //////

func (l *linkedList) insert(val int) {
  n := node{data: val}
  n.next = l.head
  l.head = &n

  l.length++
}

func (l linkedList) isEmpty() bool {
  return l.head == nil
}

func (l *linkedList) pop() int {
  l.length--
  poppedNode := l.head
  l.head = poppedNode.next 
  poppedNode.next = nil

  return poppedNode.data
}


func main() {
  // list := linkedList{}
  
  // fmt.Printf("Right now, our list looks like %T\n", list)

  // list.insert(1)
  // list.insert(2)
  // list.insert(3)
  // list.insert(4)
  // list.insert(5)

  // fmt.Printf("The head node has a value %v\n", list.head.data)

  // current := list.head

  // for current.next != nil {
  //   current = current.next
  // }

  // fmt.Printf("The tail of the linked list is %v\n", current.data)

  // fmt.Printf("The head node has a value %v\n", list.head.data)
  // fmt.Printf("Right now our list has a length of %v\n", list.length)
  // fmt.Printf("When I pop the list, the value we get is %v\n", list.pop())

  // myQueue := queue{}
  // myQueue.push(5)
  // myQueue.push(6)
  
  // myQueue.peek()
  // list := linkedList{}

  // list.insert(5)
    myQueue := newQueue()
    myQueue.push(1)
    myQueue.push(2)
    myQueue.push(3)
    fmt.Printf("The last value is %v\n", myQueue.pop())
    // fmt.Printf("The value is %#v\n", myQueue.inStack.head)
    
  // myQueue.push(1)
  // myQueue.push(2)
  // myQueue.push(3)

  // fmt.Printf("The top of the ")
}
