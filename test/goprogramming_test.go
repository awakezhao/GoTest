package test

import (
	"fmt"
	"testing"
)

type Celsius float64
type Fahrenheit float64

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

type Person struct {
	name string
}

func Test6(t *testing.T) {
	// var x int
	// x = 1
	// p := new(bool)
	// *p = true
	// person := new(Person)
	// person.name = "bob"

}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g ℃", c)
}

func Test7(t *testing.T) {
	var c Celsius = FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%g\n", c)
}
