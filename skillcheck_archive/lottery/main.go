package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// bの読み込み
	b := nextLine()

	// nの読み込み
	n, e := strconv.Atoi(nextLine())
	if e != nil {
		panic(e)
	}

	// aの読み込み
	// var numbers []string
	for range make([]int, n) {
		a := nextLine()
		if a == b {
			fmt.Println("first")
		} else if isAdjacent(a, b) {
			fmt.Println("adjacent")
		} else if equalLastDigits(a, b, 4) {
			fmt.Println("second")
		} else if equalLastDigits(a, b, 3) {
			fmt.Println("third")
		} else {
			fmt.Println("blank")
		}
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func equalLastDigits(a, b string, digits int) bool {
	a_last := a[6-digits:]
	b_last := b[6-digits:]
	if a_last == b_last {
		return true
	} else {
		return false
	}
}

func isAdjacent(a, b string) bool {
	a_int, e := strconv.Atoi(a)
	if e != nil {
		panic(e)
	}
	b_int, e := strconv.Atoi(b)
	if e != nil {
		panic(e)
	}
	if a_int+1 == b_int {
		return true
	} else if a_int-1 == b_int {
		return true
	} else {
		return false
	}
}
