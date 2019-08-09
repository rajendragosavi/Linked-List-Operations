package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
}

type Stack struct {
	nodes []*Node
	count int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func Newstack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	fmt.Println("count is :", s.count)
	s.count--
	return s.nodes[s.count]
}

// ######################## Queue Implementation ######################

/*
	tail ->  | | | | | | | head ->
	for implementation we will used fixed size queue.
	we keep track of length and capacity of the queue,
	When we insert an element at tail we increase the index for tail, coversly when we are remove/POP/Dequeue
	an element from queue, we increamene the head value .

*/

type Queue struct {
	nodes  []*Node
	cap    int
	length int
	head   int
	tail   int
}

//NewQueue function will return
func NewQueue(cap int) *Queue {
	return &Queue{
		nodes:  make([]*Node, cap),
		cap:    cap,
		length: 0,
		head:   0,
		tail:   0,
	}
}

func (q *Queue) isempty() bool {
	return q.length == 0
}
func (q *Queue) isfull() bool {
	return q.length == q.cap
}

func (q *Queue) Enqueue(item *Node) error {
	if q.isfull() {
		return errors.New("Queue is already full")
	}
	// insert the new node at the tail
	q.nodes[q.tail] = item
	q.tail++
	q.length++
	return nil
}

func (q *Queue) Dequeue() *Node {

	result := q.nodes[q.head]
	q.head++
	q.length--
	return result
}

func main() {
	s := Newstack()
	s.Push(&Node{3})
	s.Push(&Node{4})
	s.Push(&Node{5})
	fmt.Println("POP Operations :", s.Pop(), s.Pop(), s.Pop())
}
