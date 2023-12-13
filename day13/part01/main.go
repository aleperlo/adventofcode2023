package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var res int
	var pattern []string

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			res += findMirror(pattern)
			pattern = []string{}
		} else {
			pattern = append(pattern, scanner.Text())
		}
	}
	res += findMirror(pattern)

	fmt.Println(res)
}

func findMirror(pattern []string) int {
	var bound int
	var horizontal, vertical bool
	var hchecked, vchecked bool

	if len(pattern) > len(pattern[0]) {
		bound = len(pattern)
	} else {
		bound = len(pattern[0])
	}

	for i := 0; i < bound; i++ {
		horizontal, vertical = true, true
		hchecked, vchecked = false, false
		for j := 0; j < bound; j++ {
			if i+1 < len(pattern) {
				hchecked = true
				if j < len(pattern[0]) && pattern[i][j] != pattern[i+1][j] {
					horizontal = false
				}
			}
			if i+1 < len(pattern[0]) {
				vchecked = true
				if j < len(pattern) && pattern[j][i] != pattern[j][i+1] {
					vertical = false
				}
			}
		}
		if horizontal && hchecked && checkHorizontal(pattern, i) {
			fmt.Println((i + 1) * 100)
			return (i + 1) * 100
		}
		if vertical && vchecked && checkVertical(pattern, i) {
			fmt.Println(i + 1)
			return i + 1
		}
	}

	return -1
}

func checkHorizontal(pattern []string, index int) bool {
	var upper int

	if len(pattern)-index-1 < index+1 {
		upper = len(pattern) - index
	} else {
		upper = index + 2
	}

	for i := 2; i < upper; i++ {
		if pattern[index+i] != pattern[index-i+1] {
			return false
		}
	}

	return true
}

func checkVertical(pattern []string, index int) bool {
	var upper int

	if len(pattern[0])-index-1 < index+1 {
		upper = len(pattern[0]) - index
	} else {
		upper = index + 2
	}

	for i := 2; i < upper; i++ {
		for j := 0; j < len(pattern); j++ {
			if pattern[j][index+i] != pattern[j][index-i+1] {
				return false
			}
		}
	}

	return true
}
