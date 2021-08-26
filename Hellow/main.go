package main

import (
	"fmt"
)

// const (
// 	a = 100
// 	b
// 	c = 101
// 	d
// )

// const (
// 	a1 = 1
// 	a2
// 	a3
// 	a4 = iota
// 	a5
// )

// const (
// 	a, b = iota + 1, iota + 2 //1,2
// 	c, d                      //2,3
// 	e, f                      //3,4
// )

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func main() {
	s1 := "白萝卜"
	s2 := []rune(s1)
	fmt.Printf("%T\n", s2)
	s2[0] = '红'
	fmt.Println(string(s2))

	// traversalString()

	// str := "123456你好啊"
	// for _, v := range str {
	// 	fmt.Printf("%c type:%T\n", v, v)
	// }

	// res := strings.Split("123,111123", ",")
	// fmt.Println(res)

	// f1 := 1.1
	// f2 := float32(2.2)
	// f3 := f1 + f2

	// f1 := 1.23456
	// fmt.Printf("%T", f1)

	// fmt.Println(a, b, c, d, e, f)

	// fmt.Println(a1, a2, a3, a4, a5)

	// fmt.Println(a, b, c, d)

	// {
	// 	name1 := "123123213"
	// 	fmt.Printf("name1: %s", name1)
	// }
	// {
	// 	name1 := "123123213"
	// 	fmt.Printf("name1: %s", name1)
	// }
}
