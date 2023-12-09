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
	var seq []int

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()
		seq = []int{}
		for _, v := range strings.Split(line, " ") {
			n, err := strconv.Atoi(v)
			if err == nil {
				seq = append(seq, n)
			}
		}
		res += nextValue(seq)
	}

	fmt.Println(res)
}

func nextValue(seq []int) int {
	var i, next int
	zeroes := false
	last := []int{}

	for !zeroes {
		zeroes = true
		last = append(last, seq[len(seq)-1-i])
		for j := 0; j < len(seq)-1-i; j++ {
			seq[j] = seq[j+1] - seq[j]
			zeroes = zeroes && seq[j] == 0
		}
		i++
	}

	for _, v := range last {
		next += v
	}

	return next
}
