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

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		res += parseLine(scanner.Text())
	}

	fmt.Println(res)
}

func parseLine(line string) int {
	var n int
	res := 1
	nums := map[int]bool{}

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
			n++
			res *= 2
		}
	}

	if n == 0 {
		return 0
	}
	return res / 2
}
