package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func main() {
	var res int
	var directions, pos string
	var tmp Node
	nodes := map[string]Node{}

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	scanner.Scan()
	directions = scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s = %s %s", &pos, &tmp.left, &tmp.right)
		tmp.left = strings.Trim(tmp.left, "(),")
		tmp.right = strings.Trim(tmp.right, "(),")
		nodes[pos] = tmp
	}

	pos = "AAA"
	for pos != "ZZZ" {
		for _, dir := range directions {
			switch dir {
			case 'L':
				pos = nodes[pos].left
			case 'R':
				pos = nodes[pos].right
			}
			res += 1
			if pos == "ZZZ" {
				fmt.Println(res)
				return
			}
		}
	}

}
