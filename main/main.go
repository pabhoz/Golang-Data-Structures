package main

import (
	"bufio"
	"fmt"
	"os"

	binarytree "../libs/BinaryTree"
	dll "../libs/DLL"
	dictionary "../libs/Dictionary"
	hashtable "../libs/HashTable"
	queue "../libs/Queue"
	stack "../libs/Stack"
)

func main() {
	//dlls()
	//stacks()
	//queues()
	//bts()
	//dictionaries()
	hashTables()
}

func hashTables() {
	ht := hashtable.New()
	ht.Add(9864567654, "bar")
	fmt.Println(ht)
}

func dictionaries() {
	dict := dictionary.New()
	dict.Add("foo", "bar")
	fmt.Println(dict)
}

func dlls() {
	//DLL
	dll := dll.New()
	fmt.Println(dll)
}

func stacks() {
	//Stacks
	stack := stack.New(5)
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
	fmt.Println(stack)
}

func queues() {
	//Queues
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

func bts() {
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

func io() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Recibiendo texto con buffer I/O: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Print("Recibiendo texto con scan: ")
	input := ""
	fmt.Scan(&input)
	fmt.Println(input)
}
