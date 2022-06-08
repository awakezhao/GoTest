package util

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	v := GetToday0THTime()
	fmt.Println(v)

	v = GetTomorrow0THTime()
	fmt.Println(v)

	v = GetYesterday0THTime()
	fmt.Println(v)

	lg(FormatTimeToYMD(time.Now()))
}
