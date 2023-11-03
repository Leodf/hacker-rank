package main

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	for i := 1; i <= int(n); i++ {
		fmt.Println(strings.Repeat(" ", int(n)-i) + strings.Repeat("#", i))
	}
}

func StairCase(){
	n := int32(100)
	staircase(n)
}