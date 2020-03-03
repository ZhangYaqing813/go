package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Student struct {
	name  string
	age   int
	score int
}

type Stu []Student

func (s Stu) Len() int {
	return len(s)
}

func (s Stu) Less(i, j int) bool {
	return s[i].score > s[j].score
}

func (s Stu) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var students Stu
	for i := 0; i < 10; i++ {
		stud := Student{
			name:  fmt.Sprintf("student%d", rand.Intn(100)),
			age:   i,
			score: rand.Intn(100),
		}
		students = append(students, stud)
	}

	fmt.Println(students)

	sort.Sort(students)
	fmt.Println()
	fmt.Println(students)
}
