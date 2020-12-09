package main

import (
	"fmt"
	"strings"
)

type janken struct {
	g, c, p int
}

func main() {
	// m、nの読み込み
	var m, n int
	fmt.Scan(&n, &m)

	// sの読み込み
	var s string
	fmt.Scan(&s)

	// 相手が出す手の数をカウントする
	s_slice := strings.Split(s, "")
	opponent := janken{0, 0, 0}
	for _, v := range s_slice {
		if v == "G" {
			opponent.g++
		} else if v == "C" {
			opponent.c++
		} else if v == "P" {
			opponent.p++
		}
	}

	// チョキ、パーの組み合わせを洗い出す
	var combination []janken
	for c := 0; c <= n; c++ {
		for p := 0; p <= n; p++ {
			if c*2+p*5 == m {
				combination = append(combination, janken{n - c - p, c, p})
			}
		}
	}

	// 一番多く勝てる組み合わせを探す
	result := 0
	for _, combi := range combination {
		win := 0

		// パーで勝てる数をカウント
		if opponent.g >= combi.p {
			win += opponent.g
		} else {
			win += combi.p
		}

		// グーで勝てる数をカウント
		if opponent.c >= combi.g {
			win += opponent.c
		} else {
			win += combi.g
		}

		// チョキで勝てる数をカウント
		if opponent.p >= combi.c {
			win += opponent.p
		} else {
			win += combi.c
		}

		// 勝利数を更新
		if result < win {
			result = win
		}
	}

	// 結果を出力
	fmt.Println(result)
}
