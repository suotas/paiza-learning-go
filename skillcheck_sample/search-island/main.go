package main

import (
	"fmt"
)

func main() {

	// m、nの読み込み
	var m, n int
	fmt.Scan(&m, &n)

	// inputの読み込み
	input := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		input[i] = make([]int, m+2)
	}
	for i := 1; i <= n; i++ {
		input[i][0] = 0
		for j := 1; j <= m; j++ {
			if j != 0 {
				fmt.Scan(&input[i][j])
			}
		}
		input[i][m+1] = 0
	}

	// マップを端から精査
	// 1が見つかった場合はチェック関数へ渡す
	result := 0
	for y, slice := range input {
		for x, v := range slice {
			if v == 1 {
				result += island_check(x, y, input)
			}
		}
	}
	fmt.Println(result)
}

// 指定されたマスの上下左右に島がないか探す
func island_check(x, y int, input [][]int) int {
	var island [][]int
	island = append(island, []int{x, y})
	for len(island) > 0 {

		x_y := island[len(island)-1]
		x := x_y[0]
		y := x_y[1]
		island = island[:len(island)-1]

		input[y][x] = 0

		// 下に繋がった島がないか確認
		if input[y+1][x] == 1 {
			island = append(island, []int{x, y + 1})
		}
		// 右に繋がった島がないか確認
		if input[y][x+1] == 1 {
			island = append(island, []int{x + 1, y})
		}
		// 上に繋がった島がないか確認
		if input[y-1][x] == 1 {
			island = append(island, []int{x, y - 1})
		}
		// 左に繋がった島がないか確認
		if input[y][x-1] == 1 {
			island = append(island, []int{x - 1, y})
		}
	}
	return 1
}
