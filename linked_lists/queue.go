package main

type queue struct {
  inStack *linkedList
  outStack *linkedList
}

// QUEUE METHODS /////

func (q *queue) push(val int) {
  q.inStack.insert(val)
}

func newQueue() queue {
  q := queue{}
  q.inStack, q.outStack = &linkedList{}, &linkedList{}

  return q
}

func (q *queue) pop() int {
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

