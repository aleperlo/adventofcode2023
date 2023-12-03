package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type pos struct {
	r int
	c int
}

func main() {
	var res int
	var grid []string
	var num string
	var adj map[pos]bool
	var gears map[pos][]int
	var p pos

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
	adj = map[pos]bool{}
	gears = map[pos][]int{}
	for i, r := range grid {
		for j, c := range r {
			p.r = i
			p.c = j
			if unicode.IsDigit(c) {
				num += string(c)
				searchGears(grid, i, j, adj)
			}
			if (!unicode.IsDigit(c) || j == len(grid[0])-1) && num != "" {
				val, err := strconv.Atoi(num)
				if err == nil {
					for k, _ := range adj {
						gears[k] = append(gears[k], val)
					}
				}
				num = ""
				adj = map[pos]bool{}
			}
		}
	}

	res = totalGearRatio(gears)
	fmt.Println(res)
}

func searchGears(grid []string, r int, c int, adj map[pos]bool) {
	var p pos
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if r+i >= 0 && r+i < len(grid) && c+j >= 0 && c+j < len(grid[0]) {
				if rune(grid[r+i][c+j]) == '*' {
					p.r = r + i
					p.c = c + j
					adj[p] = true
				}
			}
		}
	}
}

func totalGearRatio(gears map[pos][]int) int {
	var res, inc int
	for _, vals := range gears {
		if len(vals) < 2 {
			inc = 0
		} else {
			inc = 1
			for _, val := range vals {
				inc *= val
			}
		}
		res += inc
	}
	return res
}
