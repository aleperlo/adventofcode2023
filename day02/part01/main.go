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
		if validGame(g) {
			res += g.id
		}
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

func validGame(g Game) bool {
	valid := true
	for _, e := range g.extr {
		if e.r > 12 || e.g > 13 || e.b > 14 {
			valid = false
			break
		}
	}
	return valid
}
