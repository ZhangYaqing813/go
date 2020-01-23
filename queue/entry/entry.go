package main

import (
	"fmt"
	"zyq.cn/queue"
)

func main() {
	q :=queue.Queue{1}
	q.Push(4)
	q.Push(5)
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
	q.Pop()
	fmt.Println(q.IsEmpty())


}
