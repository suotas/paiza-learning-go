package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	// nの読み込み
	n := nextInt()

	// 各a_iについて、7で割った余りの出現回数を数える
	counts := make([]int, 7)
	for range make([]int, n) {
		counts[nextInt()%7]++
	}

	// 7で割り切れる組み合わせの総数
	sum := 0

	// 0-6の数字から1つ目の数字を選ぶループ
	for i := 0; i < 7; i++ {
		// 0-6の数字から2つ目の数字を選ぶループ
		for j := i; j < 7; j++ {
			// 0-6の数字から3つ目の数字を選ぶループ
			for k := j; k < 7; k++ {
				// 3つの数字の和が7で割り切れない場合、カウントしない
				if (i+j+k)%7 != 0 {
					continue
				}

				// 3つの数字の出現回数を掛け合わせる
				c1 := counts[i]
				c2 := counts[j]
				c3 := counts[k]

				// 2つの数字が同じ場合、重複を除く為
				// 片方から1を引いて積を2で割る
				// 3つの数字が同じ場合、重複を除く為
				// 片方から1を、もう片方から2を引いて積を6で割る
				if i == j {
					c2 -= 1
				}
				if k == i {
					c3 -= 1
				}
				if k == j {
					c3 -= 1
				}

				pat := c1 * c2 * c3
				if i == j && j == k {
					pat /= 6
				} else if i == j || i == k || j == k {
					pat /= 2
				}

				sum += pat
			}
		}
	}

	// 7で割り切れた組み合わせの数を出力
	fmt.Println(sum)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
