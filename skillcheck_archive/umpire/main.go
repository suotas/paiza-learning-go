package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// カウントの初期化
	counts := make(map[string]int)

	n, e := strconv.Atoi(nextLine())
	if e != nil {
		panic(e)
	}

	// 終了判定が出るまでループ
	for range make([]int, n) {
		s := nextLine()
		if s == "strike" {
			// ストライクなら
			if counts[s] == 2 {
				// 既にストライクカウントが2だったら
				// アウト
				fmt.Println("out!")
			} else {
				// ストライクカウントを加算
				counts[s]++
				fmt.Println("strike!")
			}
		} else if s == "ball" {
			// ボールなら
			if counts[s] == 3 {
				// 既にボールカウントが3だったら
				// フォアボール
				fmt.Println("fourball!")
			} else {
				// ボールカウントを加算
				counts[s]++
				fmt.Println("ball!")
			}
		}
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
