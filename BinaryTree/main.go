package main

import (
	binarytree "../libs/BinaryTree"
)

func main() {
	//DLL
	//dll := dll.New()
	//fmt.Println(dll)

	//Stacks
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

	//Queues
	/*q := queue.New(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.PrintMe()

	q.Dequeue()
	q.PrintMe()

	q.Dequeue()
	q.PrintMe()*/

	bt := binarytree.New()
	bt.Insert() //2
	bt.Insert() //3
	bt.Insert() //4
	bt.Insert() //5
	bt.Insert() //6
	bt.Insert() //7
	bt.Insert() //8
	bt.Insert() //9
	bt.Insert() //10
	bt.Insert() //11
	bt.Insert() //12
	bt.Insert() //13
	bt.Insert() //14
	bt.Insert() //15
	bt.PrintInorder()

}
