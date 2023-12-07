package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type race struct {
	duration int
	record   int
}

func main() {
	var res int
	var tmp race
	var races []race

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	// Parse first line
	scanner.Scan()
	_, rest, _ := strings.Cut(scanner.Text(), ": ")
	for _, v := range strings.Split(rest, " ") {
		d, err := strconv.Atoi(v)
		if err == nil {
			tmp.duration = d
			races = append(races, tmp)
		}
	}

	// Parse second line
	scanner.Scan()
	_, rest, _ = strings.Cut(scanner.Text(), ": ")
	i := 0
	for _, v := range strings.Split(rest, " ") {
		r, err := strconv.Atoi(v)
		if err == nil {
			races[i].record = r
			i++
		}
	}

	res = 1
	for _, r := range races {
		n := findBounds(r)
		res *= n
	}

	fmt.Println(res)
}

func findBounds(r race) int {
	var delta, lbf, ubf float64
	var lb, ub int

	delta = float64(r.duration*r.duration) - 4.0*float64(r.record)
	lbf = (float64(r.duration) - math.Sqrt(delta)) / 2
	ubf = (float64(r.duration) + math.Sqrt(delta)) / 2

	if lbf == math.Ceil(lbf) {
		lb = int(lbf + 1)
	} else {
		lb = int(math.Ceil(lbf))
	}
	if ubf == math.Floor(ubf) {
		ub = int(ubf - 1)
	} else {
		ub = int(math.Floor(ubf))
	}
	return ub - lb + 1
}
