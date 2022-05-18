package test

import "testing"

func Test(t *testing.T) {
	byteAndRune := NewByteAndRune()
	byteAndRune.ByteTest()
	byteAndRune.RuneTest()
}
