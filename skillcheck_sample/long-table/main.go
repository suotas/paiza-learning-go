package main

import (
	"fmt"
)

func main() {
	// m、nを読み込む
	var m, n int
	fmt.Scan(&n, &m)

	// 座席を表す配列seats
	seats := make([]bool, n)

	// 全ての席をtrueで初期化
	for i, _ := range seats {
		seats[i] = true
	}

	// m回ループ
	for i := 1; i <= m; i++ {
		// a_i、b_iを読み込む
		var a, b int
		fmt.Scan(&a, &b)

		// グループ全員が着席できるフラグ
		flg := true

		// グループの人数回ループ
		for j := b; j < a+b; j++ {

			// 席を示す指標k
			k := 0

			// jがnを超える場合、指標を1から数え直す
			if j > n {
				k = j - n
			} else {
				k = j
			}

			// 一つでもfalse（着席済み）がある場合
			// グループごと帰ってしまう為、着席しない
			if !seats[k-1] {
				flg = false
			}
		}

		// グループ全員が着席できる場合
		// 座席をfalseにする
		if flg {
			for j := b; j < a+b; j++ {
				// 席を示す指標k
				k := 0

				// jがnを超える場合、指標を1から数え直す
				if j > n {
					k = j - n
				} else {
					k = j
				}
				seats[k-1] = false
			}
		}
	}

	// 最後にfalseの数を数えて出力
	result := 0
	for _, v := range seats {
		if !v {
			result++
		}
	}
	fmt.Print(result)
}
