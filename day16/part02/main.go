package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Pos struct {
	r int
	c int
}

var N = Pos{-1, 0}
var S = Pos{1, 0}
var E = Pos{0, 1}
var W = Pos{0, -1}

type Beam struct {
	p Pos
	d Pos
}

type Node struct {
	val  Beam
	next *Node
}

type Queue struct {
	head *Node
	last *Node
}

func main() {
	var res, tmp int
	var grid []string

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 {
				tmp = countEnergized(grid, Beam{Pos{i, j}, S})
				if tmp > res {
					res = tmp
				}
			}
			if i == len(grid)-1 {
				tmp = countEnergized(grid, Beam{Pos{i, j}, N})
				if tmp > res {
					res = tmp
				}
			}
			if j == 0 {
				tmp = countEnergized(grid, Beam{Pos{i, j}, E})
				if tmp > res {
					res = tmp
				}
			}
			if j == len(grid)-1 {
				tmp = countEnergized(grid, Beam{Pos{i, j}, W})
				if tmp > res {
					res = tmp
				}
			}
		}
	}

	fmt.Println(res)
}

func countEnergized(grid []string, b Beam) int {
	var q Queue
	var res int
	var c rune
	var energized [][]bool
	var visited map[Beam]bool

	energized = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		energized[i] = make([]bool, len(grid[0]))
	}
	visited = map[Beam]bool{}

	q.Enqueue(b)
	for !q.Empty() {
		b = q.Pop()
		for b.p.r >= 0 && b.p.r < len(grid) && b.p.c >= 0 && b.p.c < len(grid[0]) && !visited[b] {
			energized[b.p.r][b.p.c] = true
			visited[b] = true
			c = rune(grid[b.p.r][b.p.c])
			nb, err := b.nextBeam(c)
			if err == nil {
				q.Enqueue(nb)
			}
			b.p.Move(b.d)
		}
	}

	for _, line := range energized {
		for _, v := range line {
			if v {
				res++
			}
		}
	}

	return res
}

func (b *Beam) nextBeam(c rune) (Beam, error) {
	switch c {
	case '-':
		if b.d == N || b.d == S {
			b.d = E
			return Beam{b.p, W}, nil
		}
	case '|':
		if b.d == E || b.d == W {
			b.d = N
			return Beam{b.p, S}, nil
		}
	case '/':
		switch b.d {
		case N:
			b.d = E
		case S:
			b.d = W
		case W:
			b.d = S
		case E:
			b.d = N
		}
	case '\\':
		switch b.d {
		case N:
			b.d = W
		case S:
			b.d = E
		case W:
			b.d = N
		case E:
			b.d = S
		}
	}
	return *b, errors.New("The beam has not been split.")
}

func (p *Pos) Move(p1 Pos) {
	p.r += p1.r
	p.c += p1.c
}

func (q *Queue) Enqueue(b Beam) {
	n := Node{b, nil}
	if q.head == nil {
		q.head = &n
		q.last = &n
	} else {
		q.last.next = &n
		q.last = q.last.next
	}
}

func (q *Queue) Pop() Beam {
	var n Node
	if q.head != nil {
		n = *(q.head)
		q.head = q.head.next
	}
	return n.val
}

func (q Queue) Empty() bool {
	return q.head == nil
}
