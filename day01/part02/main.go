package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var res int

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	res = 0
	for scanner.Scan() {
		res += getCalibrationValue(scanner.Text())
	}

	fmt.Println(res)
}

func getCalibrationValue(line string) int {
	var n1, n2, tmp int
	var s string

	found1 := false
	for _, c := range line {
		if unicode.IsDigit(c) {
			tmp = int(c - '0')
			s = ""
		} else {
			s = s + string(c)
			tmp = stringToInt(s)
		}
		if tmp != 0 {
			if !found1 {
				found1 = true
				n1 = tmp
			}
			n2 = tmp
		}
	}
	return (n1 * 10) + n2
}

func stringToInt(s string) int {
	n := 0
	digits := map[string]int{
		"one": 1, "two": 2, "three": 3,
		"four": 4, "five": 5, "six": 6,
		"seven": 7, "eight": 8, "nine": 9,
	}
	for k, v := range digits {
		if strings.HasSuffix(s, k) {
			n = v
			break
		}
	}
	return n
}
