package main

import (
	"bufio"
	"fmt"
	"os"
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
	var n1, n2 int
	found1 := false
	for _, v := range line {
		if unicode.IsDigit(v) {
			n2 = int(v - '0')
			if !found1 {
				found1 = true
				n1 = n2
			}
		}
	}
	return (n1 * 10) + n2
}
