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
	var s Pos
	var previous, current Pos
	var neighbours []Pos
	var countp, res int
	var pipes [][]rune
	var loop [][]bool
	var tmp rune

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	i := 0
	for scanner.Scan() {
		pipes = append(pipes, []rune{})
		for _, c := range scanner.Text() {
			pipes[i] = append(pipes[i], c)
		}
		loop = append(loop, []bool{})
		i++
	}

	for i := 0; i < len(loop); i++ {
		loop[i] = make([]bool, len(pipes[0]))
	}

	s = getStartingPos(pipes)
	tmp = rune(pipes[s.y][s.x+1])
	if tmp == '-' || tmp == 'J' || tmp == '7' {
		current = Pos{s.x + 1, s.y}
		neighbours = append(neighbours, current)
	}
	tmp = rune(pipes[s.y][s.x-1])
	if tmp == '-' || tmp == 'L' || tmp == 'F' {
		current = Pos{s.x - 1, s.y}
		neighbours = append(neighbours, current)
	}
	tmp = rune(pipes[s.y+1][s.x])
	if tmp == '|' || tmp == 'J' || tmp == 'L' {
		current = Pos{s.x, s.y + 1}
		neighbours = append(neighbours, current)
	}
	tmp = rune(pipes[s.y-1][s.x])
	if tmp == '|' || tmp == '7' || tmp == 'F' {
		current = Pos{s.x, s.y - 1}
		neighbours = append(neighbours, current)
	}

	replaceS(s, neighbours[0], neighbours[1], pipes)

	loop[current.y][current.x] = true
	previous = s
	for current != s {
		previous, current = nextPos(current, previous, pipes)
		loop[current.y][current.x] = true
	}

	tmp = '|'
	for i := 0; i < len(pipes); i++ {
		countp = 0
		tmp = '|'
		for j := 0; j < len(pipes[0]); j++ {
			if !loop[i][j] {
				if countp%2 == 1 {
					res++
				}
			} else if pipes[i][j] != '-' {
				switch pipes[i][j] {
				case '|':
					countp++
				case 'J':
					if tmp == 'F' {
						countp++
					}
				case '7':
					if tmp == 'L' {
						countp++
					}
				}
				tmp = rune(pipes[i][j])
			}
		}
		if countp%2 != 0 {
			fmt.Println(countp, i)
		}
	}

	fmt.Println(res)
}

func printLoop(loop [][]bool) {
	for i := 0; i < len(loop); i++ {
		for j := 0; j < len(loop[0]); j++ {
			if loop[i][j] {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
	}
}

func getStartingPos(pipes [][]rune) Pos {
	for i, line := range pipes {
		for j, v := range line {
			if v == 'S' {
				return Pos{j, i}
			}
		}
	}
	return Pos{-1, -1}
}

func replaceS(s Pos, first Pos, last Pos, pipes [][]rune) {
	var tmp Pos
	for _, c := range "|L-J7F" {
		pipes[s.y][s.x] = c
		_, tmp = nextPos(last, s, pipes)
		if tmp == first {
			return
		}
	}
}

func nextPos(current Pos, previous Pos, pipes [][]rune) (Pos, Pos) {
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
