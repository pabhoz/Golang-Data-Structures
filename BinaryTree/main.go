package main

import (
	queue "../libs/Queue"
)

func main() {
	//dll := dll.New()
	//fmt.Println(dll)

	/*stack := stack.New(5)
	fmt.Println(stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
	fmt.Println(stack.Top().Value)
	stack.Pop()
	fmt.Println(stack)
	fmt.Println(stack.Top().Value)
	stack.Pop()
	fmt.Println(stack)
	fmt.Println(stack.Top().Value)
	stack.Pop()
	fmt.Println(stack)*/

	q := queue.New(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.PrintMe()

	q.Dequeue()
	q.PrintMe()

	q.Dequeue()
	q.PrintMe()

}
