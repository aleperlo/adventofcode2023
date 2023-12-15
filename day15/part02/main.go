package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const boxes = 256

type Node struct {
	next   *Node
	label  string
	length int
}

type List struct {
	head *Node
}

func main() {
	var res, box, length int
	var sequences []string
	var hashmap []List
	var label, tmp string

	hashmap = make([]List, boxes)

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Scan()
	sequences = strings.Split(scanner.Text(), ",")

	for _, seq := range sequences {
		if strings.Contains(seq, "-") {
			label, _ = strings.CutSuffix(seq, "-")
			box = hash(label)
			hashmap[box].delete(label)
		} else {
			label, tmp, _ = strings.Cut(seq, "=")
			box = hash(label)
			length, _ = strconv.Atoi(tmp)
			hashmap[box].insert(label, length)
		}
	}

	for i := 0; i < boxes; i++ {
		box = i + 1
		for j, link := 1, hashmap[i].head; link != nil; j, link = j+1, link.next {
			res += box * j * link.length
		}
	}

	fmt.Println(res)
}

func hash(label string) int {
	var current int
	for _, c := range label {
		current += int(c)
		current *= 17
		current %= 256
	}
	return current
}

func (l *List) insert(label string, length int) {
	var n Node = Node{nil, label, length}
	var last *Node
	if l.head == nil {
		l.head = &n
		return
	}
	for link := l.head; link != nil; link = link.next {
		if link.label == n.label {
			link.length = n.length
			return
		}
		last = link
	}
	last.next = &n
}

func (l *List) delete(label string) {
	var prev *Node
	for link := l.head; link != nil; prev, link = link, link.next {
		if link.label == label {
			if prev == nil {
				l.head = l.head.next
				return
			}
			prev.next = link.next
			return
		}
	}
}
