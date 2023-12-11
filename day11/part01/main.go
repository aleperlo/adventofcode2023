package main

import (
	"bufio"
	"fmt"
	"os"
)

type Item struct {
	i int
	j int
	d int
}

type Node struct {
	val  Item
	next *Node
}

type Queue struct {
	head *Node
	last *Node
}

func main() {
	var ng, nr, res int
	emptyCols, emptyRows := map[int]bool{}, map[int]bool{}
	var sky [][]int
	var galaxies []Item
	var anyGalaxies bool

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		sky = append(sky, []int{})
		anyGalaxies = false
		for _, c := range scanner.Text() {
			switch c {
			case '.':
				sky[nr] = append(sky[nr], -1)
			case '#':
				sky[nr] = append(sky[nr], ng)
				anyGalaxies = true
				ng++
			}
		}
		if !anyGalaxies {
			emptyRows[nr] = true
		}
		nr++
	}

	for j := 0; j < len(sky[0]); j++ {
		anyGalaxies = false
		for i := 0; i < len(sky); i++ {
			if sky[i][j] != -1 {
				anyGalaxies = true
			}
		}
		if !anyGalaxies {
			emptyCols[j] = true
		}
	}

	galaxies = make([]Item, ng)

	examined := 1
	for i, line := range sky {
		for j, v := range line {
			if v != -1 {
				bfs(sky, i, j, galaxies, emptyRows, emptyCols)
				for k := examined; k < ng; k++ {
					res += galaxies[k].d
				}
				examined++
			}
		}
	}

	fmt.Println(res)
}

func bfs(sky [][]int, r int, c int, galaxies []Item, emptyRows map[int]bool, emptyCols map[int]bool) {
	var q Queue
	var pos Item
	var visited [][]bool
	var inc int
	var ok bool
	width := 2

	visited = make([][]bool, len(sky))
	for i := 0; i < len(sky); i++ {
		visited[i] = make([]bool, len(sky[0]))
	}
	q.Enqueue(r, c, 0)
	visited[r][c] = true

	for !q.Empty() {
		pos = q.Pop()
		for _, d := range []int{-1, 1} {
			if pos.i+d >= 0 && pos.i+d < len(sky) && !visited[pos.i+d][pos.j] {
				_, ok = emptyRows[pos.i]
				if ok {
					inc = width
				} else {
					inc = 1
				}
				q.Enqueue(pos.i+d, pos.j, pos.d+inc)
				visited[pos.i+d][pos.j] = true
			}
			if pos.j+d >= 0 && pos.j+d < len(sky[0]) && !visited[pos.i][pos.j+d] {
				_, ok = emptyCols[pos.j]
				if ok {
					inc = width
				} else {
					inc = 1
				}
				q.Enqueue(pos.i, pos.j+d, pos.d+inc)
				visited[pos.i][pos.j+d] = true
			}
		}
		if sky[pos.i][pos.j] != -1 {
			galaxies[sky[pos.i][pos.j]] = pos
		}
	}
}

func (q *Queue) Enqueue(i int, j int, d int) {
	n := Node{Item{i, j, d}, nil}
	if q.head == nil {
		q.head = &n
		q.last = &n
	} else {
		q.last.next = &n
		q.last = q.last.next
	}
}

func (q *Queue) Pop() Item {
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
