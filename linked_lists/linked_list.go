package main

type linkedList struct {
  head *node 
  length int 
}


func (l *linkedList) insert(val int) {
  n := node{data: val, next: l.head}
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