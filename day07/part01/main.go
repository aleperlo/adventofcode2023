package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

const NKinds int = 12
const NTypes int = 7
const NCards int = 5

type Hand struct {
	cards string
	htype HandType
	bid   int
}

type Hands []Hand

func (hs Hands) Len() int {
	return len(hs)
}

func (hs Hands) Swap(i int, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func (hs Hands) Less(i int, j int) bool {
	c1, c2 := hs[i].cards, hs[j].cards
	letterValues := map[rune]int{'T': 0, 'J': 1, 'Q': 2, 'K': 3, 'A': 4}
	for c := 0; c < NCards; c++ {
		r1, r2 := rune(c1[c]), rune(c2[c])
		if r1 != r2 {
			if (unicode.IsDigit(r1) && unicode.IsDigit(r2)) || (unicode.IsDigit(r1) != unicode.IsDigit(r2)) {
				return r1 < r2
			} else {
				return letterValues[r1] < letterValues[r2]
			}
		}
	}
	return false
}

func main() {
	var res int
	hands := make([][]Hand, NTypes)
	var h Hand

	fd, err := os.Open("./../input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %d", &h.cards, &h.bid)
		h.htype = getHandType(h.cards)
		hands[h.htype-1] = append(hands[h.htype-1], h)
	}

	offset := 0
	for i := 0; i < NTypes; i++ {
		sort.Sort(Hands(hands[i]))
		for j := 0; j < len(hands[i]); j++ {
			index := offset + j + 1
			res += index * hands[i][j].bid
		}
		offset += len(hands[i])
	}

	fmt.Println(res)
}

func getHandType(cards string) HandType {
	kinds := map[rune]int{}
	types := make([]int, NCards+1)

	for _, c := range cards {
		_, ok := kinds[c]
		if ok {
			kinds[c] += 1
		} else {
			kinds[c] = 1
		}
	}

	for _, v := range kinds {
		types[v] += 1
	}

	switch {
	case types[5] == 1:
		return FiveOfAKind
	case types[4] == 1:
		return FourOfAKind
	case types[3] == 1 && types[2] == 1:
		return FullHouse
	case types[3] == 1 && types[1] == 2:
		return ThreeOfAKind
	case types[2] == 2:
		return TwoPair
	case types[2] == 1:
		return OnePair
	default:
		return HighCard
	}
}
