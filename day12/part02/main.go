package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var counter int

func main() {
	var res int
	var springs, tmps string
	var groups, tmpg []int

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		tmps, tmpg = parseLine(scanner.Text())
		springs = ""
		groups = []int{}
		n := 5
		for i := 0; i < n; i++ {
			springs += tmps
			groups = append(groups, tmpg...)
			if i != n-1 {
				springs += "?"
			}
		}
		res += countPlacements(springs, groups)
	}

	fmt.Println(res)
}

type Item struct {
	s  string
	ng uint8
	ni uint8
}

type IntItem struct {
	prev rune
	d    int
	n    int
}

func countPlacements(springs string, groups []int) int {
	var res, giTot, g, d int
	cache := map[Item]int{}
	intCache := map[IntItem]int{}

	allInt := true
	for _, v := range springs {
		if rune(v) != '?' {
			allInt = false
		}
	}

	if allInt {
		for _, v := range groups {
			giTot += v
		}
		d = len(springs) - giTot
		g = len(groups)
		res = countInt(' ', d, g, intCache)
	} else {
		res = countPlacementsRec(springs, groups, 0, 0, 0, false, cache)
	}
	return res
}

func countInt(prev rune, d int, g int, cache map[IntItem]int) int {
	var res int
	if d == 0 {
		if g <= 1 {
			return 1
		}
		return 0
	}
	if d >= 1 {
		if inc, ok := cache[IntItem{'.', d - 1, g}]; ok {
			res += inc
		} else {
			res += countInt('.', d-1, g, cache)
		}
	}
	if (prev != '#' && g >= 1) || prev == ' ' {
		if inc, ok := cache[IntItem{'#', d, g - 1}]; ok {
			res += inc
		} else {
			res += countInt('#', d, g-1, cache)
		}
	}
	cache[IntItem{prev, d, g}] = res
	return res
}

func countPlacementsRec(springs string, groups []int, ns uint8, ng uint8, ni uint8, seq bool, cache map[Item]int) int {
	var res int
	var tmp string

	if ns > 0 {
		switch springs[ns-1] {
		case '#':
			ni++
			seq = true
			if int(ng) == len(groups) || int(ni) > groups[ng] {
				return 0
			}
		case '.':
			if seq {
				if int(ng) < len(groups) && int(ni) != groups[ng] {
					return 0
				}
				ng++
			}
			seq = false
			ni = 0
		}
	}

	if int(ns) == len(springs) {
		if int(ng) < len(groups) && int(ni) == groups[ng] {
			ng++
		}
		if int(ng) == len(groups) {
			return 1
		}
		return 0
	}

	if springs[ns] == '?' {
		tmp = strings.Replace(springs, "?", ".", 1)
		if inc, ok := cache[Item{tmp[ns:], ng, ni}]; ok {
			res += inc
			counter++
		} else {
			res += countPlacementsRec(tmp, groups, ns+1, ng, ni, seq, cache)
		}
		tmp = strings.Replace(springs, "?", "#", 1)
		if inc, ok := cache[Item{tmp[ns:], ng, ni}]; ok {
			res += inc
			counter++
		} else {
			res += countPlacementsRec(tmp, groups, ns+1, ng, ni, seq, cache)
		}
	} else {
		if inc, ok := cache[Item{springs[ns:], ng, ni}]; ok {
			res += inc
			counter++
		} else {
			res += countPlacementsRec(springs, groups, ns+1, ng, ni, seq, cache)
		}
	}

	cache[Item{springs[ns:], ng, ni}] = res
	return res
}

func parseLine(line string) (string, []int) {
	var springs string = ""
	var groups []int = []int{}

	springs, tail, _ := strings.Cut(line, " ")
	for _, g := range strings.Split(tail, ",") {
		tmp, err := strconv.Atoi(g)
		if err == nil {
			groups = append(groups, tmp)
		}
	}

	return springs, groups
}
