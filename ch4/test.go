package main

import (
	"fmt"
	"strings"
)

func main() {
	for _, i := range strings.Fields("test test1 test2") {
		fmt.Println(i)
	}
}
