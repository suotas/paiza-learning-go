package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// inputの読み込み
	sc.Scan()
	input := sc.Text()

	// inputを配列にする
	arr := strings.Split(input, " ")

	// 出現回数の初期化
	counts := make(map[string]int)

	// 出現順の初期化
	var words []string

	// wordのカウント
	for _, v := range arr {
		if !contains(words, v) {
			words = append(words, v)
		}
		counts[v]++
	}

	// カウント結果の出力
	for _, v := range words {
		fmt.Printf("%s %d\n", v, counts[v])
	}
}

func contains(words []string, word string) bool {
	for _, v := range words {
		if v == word {
			return true
		}
	}
	return false
}
