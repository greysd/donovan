package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	st := s[:n%3]
	for i := 0; i < n-2; i = i + 3 {
		st = st + "," + s[n%3+i:n%3+i+3]
	}
	return strings.TrimPrefix(st, ",")
}

func main() {
	fmt.Println(comma("1234567890123456789"))
}
