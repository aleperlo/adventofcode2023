package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Range struct {
	dst int
	src int
	len int
}

type Node struct {
	data  Range
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func main() {
	var res, nSeeds, nMaps int
	var maps []BST
	var seeds, partialRes []int
	var r Range
	var wg sync.WaitGroup

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Scan()
	_, rest, _ := strings.Cut(scanner.Text(), ": ")
	for _, v := range strings.Split(rest, " ") {
		nv, err := strconv.Atoi(v)
		if err == nil {
			seeds = append(seeds, nv)
			nSeeds++
		}
	}

	partialRes = make([]int, len(seeds)/2)

	nMaps = -1
	for scanner.Scan() {
		if scanner.Text() == "" {
			scanner.Scan()
			maps = append(maps, BST{})
			nMaps++
			continue
		}
		fmt.Sscanf(scanner.Text(), "%d %d %d", &r.dst, &r.src, &r.len)
		maps[nMaps].Insert(r)
	}

	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go searchRange(seeds, i, maps, &wg, partialRes)
	}
	wg.Wait()

	res = partialRes[0]
	for i := 0; i < len(partialRes); i++ {
		if partialRes[i] < res {
			res = partialRes[i]
		}
	}

	fmt.Println(res)
}

func searchRange(seeds []int, i int, maps []BST, wg *sync.WaitGroup, res []int) {
	var location int

	defer (*wg).Done()
	res[i/2] = math.MaxInt
	for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
		location = findLocation(seed, maps)
		if location < res[i/2] {
			res[i/2] = location
		}
	}
}

func findLocation(seed int, maps []BST) int {
	var tmp int
	for _, m := range maps {
		tmp = m.Search(seed)
		if tmp != -1 {
			seed = tmp
		}
	}
	return seed
}

func (bst *BST) Insert(val Range) {
	bst.InsertRec(bst.root, val)
}

func (bst *BST) InsertRec(node *Node, val Range) *Node {
	if bst.root == nil {
		bst.root = &Node{val, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{val, nil, nil}
	}
	if val.src <= node.data.src {
		node.left = bst.InsertRec(node.left, val)
	}
	if val.src > node.data.src {
		node.right = bst.InsertRec(node.right, val)
	}
	return node
}

func (bst *BST) Search(val int) int {
	found := bst.SearchRec(bst.root, val)
	return found
}

func (bst *BST) SearchRec(node *Node, val int) int {
	if node == nil {
		return -1
	}
	if node.data.src <= val && val < (node.data.src+node.data.len) {
		return val - node.data.src + node.data.dst
	}
	if val < node.data.src {
		return bst.SearchRec(node.left, val)
	}
	if val > node.data.src {
		return bst.SearchRec(node.right, val)
	}
	return -1
}
