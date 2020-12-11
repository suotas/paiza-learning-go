package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	// Nの読み込み
	n, e := strconv.Atoi(nextLine())
	if e != nil {
		panic(e)
	}

	// Wの読み込み
	var words []string
	for range make([]int, n) {
		w := nextLine()
		for i, v := range words {
			if v == w {
				// 配列中の文字列を削除する
				words = remove(words, i)
			}
		}
		// 配列の先頭に追加する
		words = append(words, w)
	}

	// wordsを逆順で表示する
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Println(words[i])
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func remove(words []string, i int) []string {
	return append(words[:i], words[i+1:]...)
}
