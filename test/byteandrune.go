package test

import (
	"fmt"
	"unsafe"
)

type ByteAndRune struct{}

func NewByteAndRune() *ByteAndRune {
	return &ByteAndRune{}
}

func (br *ByteAndRune) ByteTest() {
	var a byte = 65
	var b uint8 = 66
	fmt.Printf("a 的值：%c \nb 的值：%c\n", a, b)
}

func (br *ByteAndRune) RuneTest() {
	var a byte = 'A'
	var b rune = 'B'
	fmt.Printf("a 占用%d 个字节数\nb 占用%d 个字节数\n", unsafe.Sizeof(a), unsafe.Sizeof(b))
}
