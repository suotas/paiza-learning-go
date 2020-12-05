package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	input := []int{a, b, c, d, e}
	min := 100
	for _, v := range input {
		if min > v {
			min = v
		}
	}
	fmt.Print(min)
}
