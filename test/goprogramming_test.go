package test

import (
	"fmt"
	"testing"
)

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
