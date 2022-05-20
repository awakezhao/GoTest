package test

import (
	"fmt"
	"testing"
)

// 常量: 一般全大写
func Test1(t *testing.T) {
	const NAME string = "kk"

	// 省略类型
	const PI = 3.1415926
	// 定义多个常量(类型相同)
	const c1, c2 int = 1, 2
	// 定义多个常量（类型不同）
	const (
		C3 string = "silence"
		C4 int    = 1
	)
	// 定义多个常量 省略类型
	const C5, C6 = 1, "silence"

	fmt.Println(NAME)
	fmt.Println(c1, c2)
	fmt.Println(C3, C4)
	fmt.Println(C5, C6)

	// 如果常量的类型和值和之前的一致，那么后面的常量的类型和值可以（必须）一起省略掉。
	const (
		C7 int = 1
		C8
		C9
		C10 float64 = 3.14
		C11
		C12
		C13 string = "kk"
	)
	fmt.Println(C7, C8, C9, C10, C11, C12, C13)

	// 枚举， const+iota
	const (
		E1 int = iota
		E2
		E3
	)
	fmt.Println(E1, E2, E3)

	const (
		E4 = (iota + 1) * 100
		E5
		E6
	)
	fmt.Println(E4, E5, E6)
}

// Test2: 作用域
func Test2(t *testing.T) {
	// 作用域： 定义标识符可以使用的范围
	// 在Go中用{}来定义作用域的范围
	// 使用原则：子语句块可以使用父语句块的标识符，父不能用子的

	outer := 1
	{
		// outer := 10
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
		outer := 21
		{
			inner2 := 3
			fmt.Println(outer, inner, inner2)
		}
	}
	// fmt.Println(inner)
}

const boilingF = 212.0
const fix = 5.0 / 9

func Test3(t *testing.T) {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g℉ or %g℃\n", f, c)
}

func Test4(t *testing.T) {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Println(fToC(freezingF))
	fmt.Println(fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * fix
}

func Test5(t *testing.T) {
	fmt.Println(f() == f())

	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
