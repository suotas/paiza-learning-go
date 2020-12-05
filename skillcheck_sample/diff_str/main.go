package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if a == b {
		fmt.Print("OK")
	} else {
		fmt.Print("NG")
	}
}
