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
	var count, res int
	var directions, pos string
	var positions []string
	var tmp Node
	var steps []int
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
		if pos[2] == 'A' {
			positions = append(positions, pos)
			steps = append(steps, 0)
		}
	}

	reached := false
	for !reached {
		for _, dir := range directions {
			count++
			reached = move(positions, dir, nodes, count, steps)
			if reached {
				break
			}
		}
	}

	res = 1
	for _, v := range steps {
		res = lcm(res, v)
	}
	fmt.Println(res)
}

func move(positions []string, dir rune, nodes map[string]Node, count int, steps []int) bool {
	reached := true
	for i, pos := range positions {
		switch dir {
		case 'L':
			positions[i] = nodes[pos].left
		case 'R':
			positions[i] = nodes[pos].right
		}
		if positions[i][2] == 'Z' {
			steps[i] = count
		}
		reached = reached && steps[i] != 0
	}
	return reached
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	if a > b {
		return a / gcd(a, b) * b
	}
	return b / gcd(a, b) * a
}
