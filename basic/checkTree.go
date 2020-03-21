package main

import "fmt"

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func (this *Tree)MakeTree(value int, pointer ...int)  {
	this.Value = value
	if cap(pointer) == 1 {
		this.Left = &Tree{Value:pointer[0]}
	}
	if cap(pointer) == 2 {
		this.Left = &Tree{Value:pointer[0]}
		this.Right = &Tree{Value:pointer[1]}
	}
}


func (this *Tree) String() string {
	if this == nil {
		return ""
	}
	s := ""
	if this.Left != nil {
		s += this.Left.String() + " "
	}
	s += fmt.Sprint(this.Value)
	if this.Right != nil {
		s += " " + this.Right.String()
	}
	return "" + s + ""
}

func main() {
	var treeA Tree
	treeA.MakeTree(3, 1, 8)
	treeA.Left.MakeTree(1, 1, 2)
	treeA.Right.MakeTree(8, 5, 13)
	ch1 := make(chan int)

	var treeB Tree
	treeB.MakeTree(8, 3, 13)
	treeB.Left.MakeTree(3, 1, 5)
	treeB.Left.Left.MakeTree(1, 1, 2)
	ch2 := make(chan int)

	go Walk(&treeA, ch1)
	go Walk(&treeB, ch2)

	for a := range ch1{
		if a != <-ch2 {
			fmt.Println("check Res", false)
		}
	}
	fmt.Println(treeA.String())

	fmt.Println(treeB.String())
}

func rangeTree(tree *Tree, ch chan int)  {
	if tree == nil {
		return
	}
	if tree.Left != nil {
		rangeTree(tree.Left, ch)
	}
	ch<-tree.Value
	if tree.Right != nil {
		rangeTree(tree.Right, ch)
	}
}

func Walk(tree *Tree, ch chan int)  {
	rangeTree(tree, ch)
	close(ch)
}
