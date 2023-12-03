package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var res int
	var grid []string
	var num string
	var adj bool

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	num = ""
	adj = false
	for i, r := range grid {
		for j, c := range r {
			if unicode.IsDigit(c) {
				num += string(c)
				adj = adj || adjacentSymbol(grid, i, j)
			}
			if adj && (!unicode.IsDigit(c) || j == len(grid[0])-1) {
				inc, err := strconv.Atoi(num)
				if err == nil {
					res += inc
				}
			}
			if !unicode.IsDigit(c) || j == len(grid[0])-1 {
				num = ""
				adj = false
			}
		}
	}

	fmt.Println(res)
}

func adjacentSymbol(grid []string, r int, c int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if r+i >= 0 && r+i < len(grid) && c+j >= 0 && c+j < len(grid[0]) {
				char := rune(grid[r+i][c+j])
				if !unicode.IsDigit(char) && char != '.' {
					return true
				}
			}
		}
	}
	return false
}
