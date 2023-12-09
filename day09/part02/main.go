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
		res += prevValue(seq)
	}

	fmt.Println(res)
}

func prevValue(seq []int) int {
	var i, prev int
	zeroes := false
	first := []int{}

	for !zeroes {
		zeroes = true
		first = append(first, seq[0])
		for j := 0; j < len(seq)-1-i; j++ {
			seq[j] = seq[j+1] - seq[j]
			zeroes = zeroes && seq[j] == 0
		}
		i++
	}

	for i := len(first) - 1; i >= 0; i-- {
		prev = first[i] - prev
	}

	return prev
}
