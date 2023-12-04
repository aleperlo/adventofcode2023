package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nn = 25

func main() {
	var res int
	var next []int
	next = make([]int, nn)

	for i := 0; i < len(next); i++ {
		next[i] = 1
	}

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		res += parseLine(scanner.Text(), next)
	}

	fmt.Println(res)
}

func parseLine(line string, next []int) int {
	var n int
	amount := next[0]
	nums := map[int]bool{}

	for i := 0; i < nn-1; i++ {
		next[i] = next[i+1]
	}
	next[nn-1] = 1

	_, rest, _ := strings.Cut(line, ":")
	p1, p2, _ := strings.Cut(rest, " | ")
	for _, v := range strings.Split(p1, " ") {
		tmp, err := strconv.Atoi(v)
		if err == nil {
			nums[tmp] = true
		}
	}

	for _, v := range strings.Split(p2, " ") {
		tmp, err := strconv.Atoi(v)
		if err == nil && nums[tmp] {
			next[n] += amount
			n++
		}
	}

	return amount
}
