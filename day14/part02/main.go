package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var res int
	var platform [][]rune
	var cache [][][]rune

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		platform = append(platform, []rune(scanner.Text()))
	}

	n := 0
	np := -1
	for np == -1 {
		cache = append(cache, platformCopy(platform))
		spin(platform)
		np = platformEqual(cache, platform)
		n++
	}
	cache = append(cache, platformCopy(platform))

	res = load(cache[np+(1000000000-np)%(n-np)])
	fmt.Println(res)
}

func platformEqual(cache [][][]rune, p2 [][]rune) int {
	var equal bool
	for np, p1 := range cache {
		equal = true
		for i := 0; i < len(p1) && equal; i++ {
			for j := 0; j < len(p1[0]) && equal; j++ {
				if p1[i][j] != p2[i][j] {
					equal = false
				}
			}
		}
		if equal {
			return np
		}
	}
	return -1
}

func platformCopy(platform [][]rune) [][]rune {
	var p [][]rune
	p = make([][]rune, len(platform))
	for i, row := range platform {
		p[i] = make([]rune, len(platform[0]))
		copy(p[i], row)
	}
	return p
}

func printPlatform(platform [][]rune) {
	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[0]); j++ {
			fmt.Print(string(platform[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func spin(platform [][]rune) {
	outer_lower := []int{0, 0, len(platform) - 1, len(platform) - 1}
	outer_upper := []int{len(platform), len(platform), -1, -1}
	inner_lower := []int{0, 0, 0, len(platform[0]) - 1}
	inner_upper := []int{len(platform[0]), len(platform[0]), len(platform[0]), -1}
	outer_inc := []int{1, 1, -1, -1}
	inner_inc := []int{1, 1, 1, -1}
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}

	for d := 0; d < 4; d++ {
		for i := outer_lower[d]; i != outer_upper[d]; i += outer_inc[d] {
			for j := inner_lower[d]; j != inner_upper[d]; j += inner_inc[d] {
				if platform[i][j] == 'O' {
					move(platform, i, j, dr[d], dc[d])
				}
			}
		}
	}
}

func move(platform [][]rune, r int, c int, dr int, dc int) {
	for i, j := r+dr, c+dc; i >= 0 && i < len(platform) && j >= 0 && j < len(platform[0]) && platform[i][j] != '#' && platform[i][j] != 'O'; i, j = i+dr, j+dc {
		platform[r][c] = '.'
		platform[i][j] = 'O'
		r = i
		c = j
	}
}

func load(platform [][]rune) int {
	var res int

	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform); j++ {
			if platform[i][j] == 'O' {
				res += len(platform) - i
			}
		}
	}

	return res
}
