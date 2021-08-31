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

//a : 1 b:2 c:3
func f1(a, b int, c bool) int {
	if c {
		return a + b
	} else {
		return a - b
	}
}

/*
	123
	123
	123
*/
func f2() {
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")
}

func f3() int {
	x := 5
	defer func(aa int) {
		aa++
	}(x)
	return x
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a, b := 1, 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	//fmt.Println(f3())

	// fmt.Println(f1(1, 2, false))

	// n := 18
	// fmt.Println(&n)
	// var a1 = [...]int{3,7,8,9,1}
	// sort.Ints(a1[:])
	// fmt.Println(a1)

	// s1 := []string{"A", "B"}
	// fmt.Println(len(s1), cap(s1))
	// s1 = append(s1, "C")
	// fmt.Println(len(s1), cap(s1))
	// s2 := append(s1, "D")
	// fmt.Println(len(s1), cap(s1))
	// fmt.Println(len(s2), cap(s2))

	// a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	// s3 := a1[:4]
	// fmt.Printf("%T", s3)

	// var s1 = []string{"A", "B", "C"}
	// fmt.Printf("type:%T len(s1):%d cap(s1):%d", s1, len(s1), cap(s1))

	// var s []int
	// fmt.Printf("%T\n", s)
	// fmt.Println(s == nil)

	// var a1 [3]int32
	// var a2 [4]bool
	// fmt.Printf("a1:%T a2:%T", a1, a2)

	// flag := "我是中国人"
	// var a int32
	// a = int32(flag)
	// fmt.Println(a)

	// s1 := "白萝卜"
	// s2 := []rune(s1)
	// fmt.Printf("%T\n", s2)
	// s2[0] = '红'
	// fmt.Println(string(s2))

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
