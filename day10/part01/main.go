package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x int
	y int
}

func main() {
	var p Pos
	var previous, current []Pos
	var res int
	var pipes []string
	var tmp rune
	var ok bool

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		pipes = append(pipes, scanner.Text())
	}

	p = getStartingPos(pipes)
	tmp = rune(pipes[p.y][p.x+1])
	if tmp == '-' || tmp == 'J' || tmp == '7' {
		current = append(current, Pos{p.x + 1, p.y})
		previous = append(previous, p)
	}
	tmp = rune(pipes[p.y][p.x-1])
	if tmp == '-' || tmp == 'L' || tmp == 'F' {
		current = append(current, Pos{p.x - 1, p.y})
		previous = append(previous, p)
	}
	tmp = rune(pipes[p.y+1][p.x])
	if tmp == '|' || tmp == 'J' || tmp == 'L' {
		current = append(current, Pos{p.x, p.y + 1})
		previous = append(previous, p)
	}
	tmp = rune(pipes[p.y-1][p.x])
	if tmp == '|' || tmp == '7' || tmp == 'F' {
		current = append(current, Pos{p.x, p.y - 1})
		previous = append(previous, p)
	}

	ok = false
	res = 1
	for !ok {
		ok = true
		for i := 0; i < len(current); i++ {
			previous[i], current[i] = nextPos(current[i], previous[i], pipes)
		}
		ok = ok && current[0] == current[1]
		res++
	}

	fmt.Println(res)
}

func getStartingPos(pipes []string) Pos {
	for i, line := range pipes {
		for j, v := range line {
			if v == 'S' {
				return Pos{j, i}
			}
		}
	}
	return Pos{-1, -1}
}

func nextPos(current Pos, previous Pos, pipes []string) (Pos, Pos) {
	var diff, next Pos
	var pipe rune
	diff.x = current.x - previous.x
	diff.y = current.y - previous.y
	pipe = rune(pipes[current.y][current.x])

	next = current
	switch pipe {
	case '|':
		next.y += diff.y
	case '-':
		next.x += diff.x
	case 'L':
		if diff.y == 1 {
			next.x++
		} else {
			next.y--
		}
	case 'J':
		if diff.y == 1 {
			next.x--
		} else {
			next.y--
		}
	case '7':
		if diff.y == -1 {
			next.x--
		} else {
			next.y++
		}
	case 'F':
		if diff.y == -1 {
			next.x++
		} else {
			next.y++
		}
	}

	return current, next
}
