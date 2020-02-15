package main

import (
	"fmt"
)

type Object interface {
}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

//判断链表是否为空
func (this *List) isEmpty() bool {
	if (*this).headNode == nil {
		return true
	} else {
		return false
	}
}

//列表长度
func (this *List) Length() int {
	cur := this.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

//从头部插入元素
func (this *List) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = (*this).headNode
	this.headNode = node
	return node
}

// 从尾部添加元素
func (this *List) Append(data Object) {
	node := &Node{Data: data}
	if this.isEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for {
			if cur.Next != nil {
				cur = cur.Next
			} else {
				cur.Next = node
				break
			}
		}
	}

}

//从指定位置添加元素
func (this *List) insert(index int, data Object) {
	/*
		0: 表示在链表头田间
		>length: 在链表尾部添加
		0< index < length :在指定的位置添加
	*/

	if index == 0 {
		this.Add(data)
	} else if index >= this.Length() {
		this.Append(data)
	} else {
		count := 0
		cur := this.headNode
		for count < (index - 1) {
			cur = cur.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = cur.Next
		cur.Next = node
	}
}

// 删除指定的元素
func (this *List) remove(data Object) {
	if this.isEmpty() {
		fmt.Println("This list is empty")
	}

	cur := this.headNode
	if cur.Data == data {
		this.headNode = cur.Next
	} else {
		for cur.Next != nil {
			if cur.Data == data {
				cur.Next = cur.Next.Next
			} else {
				cur = cur.Next
			}
		}
	}
}

// 删除指定位置的元素
func (this *List) deleteIndex(index int) { //链表位置是从零开始
	//计算，比如index = 2 ,实际删除链表中的第三个元素
	cur := this.headNode

	if index == 0 {
		this.headNode = cur.Next
	}

	if index > this.Length() {
		fmt.Println("位置信息大于list 长度")
		return
	}
	count := 0
	for cur.Next != nil {
		if count != (index - 1) {
			count++
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
			break
		}

	}
}

//查找元素

func (this *List) find(data Object) bool {
	cur := this.headNode
	for cur.Next != nil {
		if cur.Data == data {
			return true
		} else {
			cur = cur.Next
		}
	}
	return false
}

//遍历所有节点
func (this *List) showList() {
	if !this.isEmpty() {
		cur := this.headNode
		for {
			if cur.Next == nil {
				fmt.Printf("\t%v", cur.Data)
				break
			} else {
				fmt.Printf("\t%v", cur.Data)
				cur = cur.Next
			}
		}
	}

}

func main() {
	list := List{}
	list.Add(1)
	list.Append(2)
	list.Append(3)
	list.Append(5)
	list.Append(6)

	list.showList()
	data := 4
	fmt.Println()
	fmt.Printf("list is length：%v\n", list.Length())
	fmt.Println()
	fmt.Printf("%v 是否存在于list 中 ？，%v ",
		data, list.find(data))

	//list.deleteIndex(2)
	//fmt.Println()
	//list.showList()
	//fmt.Printf("list %v\n", list)
}
