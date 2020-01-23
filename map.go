package main

import "fmt"

func main ()  {
	m := map[string]string{
		"name": "zhangsan",
		"course": "golang",
	}

	m2 := make(map[string]string)

	var m3 map[string]int

	fmt.Println(m2, m, m3)
	fmt.Println("======")

	for k, v := range m{
		fmt.Println(k, v)
	}

	fmt.Println("++++++++")
	couseName := m["course"]
	fmt.Println(couseName)
// 判断key 是否存在map 中
	if couseName, ok :=m["course"]; ok{
		fmt.Println(couseName)
	}else {
		fmt.Println("key does not exist")
	}
// delete 某个key

	fmt.Println("---------")
	if name, ok := m["name"]; ok{
		fmt.Println(name,ok)
		delete(m,"name")
	} else {
		fmt.Println("The key dose not exsit!")
	}
	fmt.Println("---------")
	fmt.Println("**********")
	//str := make(map[byte]int)
	s := "abcdeffed"
	for i, ch := range []byte(s) {
		//fmt.Println(str[ch],i)

		fmt.Println(i,ch)

	}

}