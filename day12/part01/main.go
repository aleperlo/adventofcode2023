package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var res int
	var springs []rune
	var groups []int

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		springs, groups = parseLine(scanner.Text())
		res += countPlacements(springs, groups)
	}

	fmt.Println(res)
}

func countPlacements(springs []rune, groups []int) int {
	return countPlacementsRec(springs, groups, 0, 0, 0, false)
}

func countPlacementsRec(springs []rune, groups []int, ns int, ng int, ni int, seq bool) int {
	var res int

	if ns > 0 {
		switch springs[ns-1] {
		case '#':
			ni++
			seq = true
			if ng == len(groups) || ni > groups[ng] {
				return 0
			}
		case '.':
			if seq {
				if ng < len(groups) && ni != groups[ng] {
					return 0
				}
				ng++
			}
			seq = false
			ni = 0
		}
	}

	if ns == len(springs) {
		if ng < len(groups) && ni == groups[ng] {
			ng++
		}
		if ng == len(groups) {
			return 1
		}
		return 0
	}

	if springs[ns] == '?' {
		springs[ns] = '.'
		res += countPlacementsRec(springs, groups, ns+1, ng, ni, seq)
		springs[ns] = '#'
		res += countPlacementsRec(springs, groups, ns+1, ng, ni, seq)
		springs[ns] = '?'
	} else {
		res += countPlacementsRec(springs, groups, ns+1, ng, ni, seq)
	}
	return res
}

func parseLine(line string) ([]rune, []int) {
	var springs []rune = []rune{}
	var groups []int = []int{}

	head, tail, _ := strings.Cut(line, " ")
	for _, s := range head {
		springs = append(springs, s)
	}
	for _, g := range strings.Split(tail, ",") {
		tmp, err := strconv.Atoi(g)
		if err == nil {
			groups = append(groups, tmp)
		}
	}

	return springs, groups
}
