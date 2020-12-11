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

	var input []int
	for range make([]int, n) {
		input = append(input, nextInt())
	}

	// カードの組み合わせを作る
	combination := getAllConbinationAt3(input)

	// 配列ごとに合計値が7で割り切れるかチェック
	sum := 0
	for _, v := range combination {
		sum += isMod7(v)
	}

	// 7で割り切れた配列の数を出力
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

func getAllConbinationAt3(card []int) [][]int {

	var result [][]int

	for i := 0; i < len(card); i++ {
		for j := i + 1; j < len(card); j++ {
			for k := j + 1; k < len(card); k++ {
				result = append(result, []int{card[i], card[j], card[k]})
			}
		}
	}

	return result
}

func isMod7(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%7 == 0 {
		return 1
	} else {
		return 0
	}
}
