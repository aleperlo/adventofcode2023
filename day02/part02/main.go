package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Extraction struct {
	r int
	g int
	b int
}

type Game struct {
	id   int
	n    int
	extr []Extraction
}

func main() {
	var res int
	var g Game

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		g = parseLine(scanner.Text())
		res += power(g)
	}

	fmt.Println(res)
}

func parseLine(line string) Game {
	var g Game
	var tmp Extraction
	var color string
	var nc int

	splitLine := strings.Split(line, ":")
	fmt.Sscanf(splitLine[0], "Game %d", &g.id)
	for i, e := range strings.Split(splitLine[1], ";") {
		g.extr = append(g.extr, tmp)
		for _, c := range strings.Split(e, ",") {
			fmt.Sscanf(c, "%d %s", &nc, &color)
			switch color {
			case "red":
				g.extr[i].r = nc
			case "green":
				g.extr[i].g = nc
			case "blue":
				g.extr[i].b = nc
			}
		}
		g.n++
	}

	return g
}

func power(g Game) int {
	var res Extraction

	for _, e := range g.extr {
		if e.r > res.r {
			res.r = e.r
		}
		if e.g > res.g {
			res.g = e.g
		}
		if e.b > res.b {
			res.b = e.b
		}
	}

	return res.r * res.g * res.b
}
