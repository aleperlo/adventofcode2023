package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var res int

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println(res)
}
