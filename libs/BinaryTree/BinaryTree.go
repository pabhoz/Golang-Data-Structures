package binarytree

import (
	"fmt"

	queue "../Queue"
)

//Tree BT structure
type Tree struct {
	left  *Tree
	Value int
	right *Tree
}

// New returns a bt
func New() *Tree {
	tree := new(Tree)
	tree.Value = 1
	return tree
}

//Insert a new node into the tree
func (t *Tree) Insert() *Tree {

	node := new(Tree)
	q := queue.New(100)
	q.Enqueue(t)

	for !q.IsEmpty() {
		temp := q.Front().Value.(*Tree) //Casting interface to *Tree
		q.Dequeue()

		if temp.left == nil {
			node.Value = temp.Value * 2
			temp.left = node
			break
		} else {
			q.Enqueue(temp.left)

			if temp.right == nil {
				node.Value = (temp.Value * 2) + 1
				temp.right = node
				break
			} else {
				q.Enqueue(temp.right)
			}
		}
	}

	return node
}

// Left returns left children
func (t *Tree) Left() *Tree {
	return t.left
}

// Right Returns rigth children
func (t *Tree) Right() *Tree {
	return t.right
}

// Level returns tree level
func (t *Tree) Level() int {
	level := 0
	previous := -1 * ((1 - t.Value) / 2)

	for previous >= 1 {
		level++
		previous = -1 * ((1 - previous) / 2)
	}
	return level
}

// PrintInorder prints tree inorder traverse
func (t *Tree) PrintInorder() {
	if t != nil {
		t.left.PrintInorder()
		fmt.Println(t.Value, " ")
		t.right.PrintInorder()
	}
}
