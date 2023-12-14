package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var res int
	var platform [][]rune

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		platform = append(platform, []rune(scanner.Text()))
	}

	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[0]); j++ {
			if platform[i][j] == 'O' {
				move(platform, i, j, -1, 0)
			}
		}
	}

	res = load(platform)
	fmt.Println(res)
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
