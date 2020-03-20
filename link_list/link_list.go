package main

import (
	"errors"
	"fmt"
)

/**
	尝试实现：
	Push()
	Pop()
	Unshift()
	Shift()
	Insert()
	Print()
	ToJson()
	Clone()
 */
type LinkNode struct {
	Data interface{}
	Next *LinkNode
	Pre *LinkNode
}
func  (this *LinkNode) IsNull() bool  {
	return this == new(LinkNode).Next
}

type LinkList struct {
	first *LinkNode
	last *LinkNode
	current *LinkNode
}

func  (this *LinkList) Push(data interface{})  {
	if nil == this.first {
		node := LinkNode{Data:data}
		this.first = &node
		this.last = &node
		return
	}
	node := LinkNode{data, nil, this.last}
	this.last.Next = &node
	this.last = &node
}

func (this *LinkList) Pop() (last LinkNode) {
	if this.IsNull() {
		panic("LinkList is nil")
	}
	last = *this.last
	this.last = this.last.Pre
	this.last.Next = nil
	return last
}

func (this *LinkList) IsNull() bool {
	return this.last == nil
}

func (this *LinkList) Next() (LinkNode, error) {
	if this.current == nil {
		this.current = this.first
		return *this.current, nil
	} else if !this.current.Next.IsNull() {
		println(this.current.Next, new(LinkNode), this.current.Next.IsNull())
		this.current = this.current.Next
		return *this.current, nil
	}
	return LinkNode{}, errors.New("到头了")
}

func (this *LinkList) Print() {
	for current,err := this.Next(); err == nil; current,err = this.Next() {
		fmt.Println(current.Data)
	}
}



func main() {
	arr := []int{1,2,3,4,5,6}
	linkList := LinkList{}
	for _,v := range arr{
		linkList.Push(v)
	}
	fmt.Println(linkList)
	fmt.Println(*linkList.first, *linkList.last)

	linkList.Print()
}
