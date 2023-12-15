package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var res, inc int
	var seq string

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Scan()
	seq = scanner.Text()

	for _, c := range seq {
		if c == ',' {
			res += inc
			inc = 0
		} else {
			inc = hash(inc, c)
		}
	}
	res += inc

	fmt.Println(res)
}

func hash(current int, c rune) int {
	current += int(c)
	current *= 17
	current %= 256
	return current
}
