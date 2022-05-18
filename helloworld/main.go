// package必须是文件的代码的第一行，不包括注释

package main

import "fmt"

/*
	多行注释
*/

var k string = "k"
var kk string = "kk"

// kkk := "kkk" //简短声明只能在函数内部使用
// fmt.Println("hello world")

var a = 10
var b = 20
var c = a + b

func main() {
	// 单行注释
	fmt.Println("hello, world")
}
