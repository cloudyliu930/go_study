package main

import "fmt"

var (
	aa = 3
	bb = 4
	cc = 5
)

func test1() {
	a := 1
	b := 2
	fmt.Println(a, b)
}

func test2() {
	fmt.Println(aa, bb, cc)
}

func main() {
	fmt.Println("hello world!!!!")

	var a int = 4
	var str string = "abc"
	fmt.Print(a, str)

	test1()
	test2()
}
