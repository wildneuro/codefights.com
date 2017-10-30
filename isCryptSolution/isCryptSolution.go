package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isCryptSolution(crypt []string, solution [][]string) bool {

	fmt.Println(crypt)
	fmt.Println(solution)

	// Replace all crypts with numbers:
	for i, s := range solution {
		fmt.Printf("[%d]:%s:%s\n", i, s[0], s[1])
		for j, _ := range crypt {
			crypt[j] = strings.Replace(crypt[j], s[0], s[1], -1)
		}
	}

	// Calc them:
	sumIn := 0
	sumRef := 0
	leadingZeros := 0
	for j, next := range crypt {
		num, _ := strconv.Atoi(next)
		fmt.Println("Next:", next, next[0] == '0', len(next))
		if next[0] == '0' && len(next) > 1 {
			leadingZeros++
		}
		if j < len(crypt)-1 {
			sumIn += num
		} else {
			sumRef = num
		}
	}

	fmt.Println(leadingZeros, len(crypt)-1)
	fmt.Println(sumIn)
	fmt.Println(sumRef)

	if sumIn != sumRef {
		return false
	}

	if leadingZeros > 0 {
		return false
	}

	return true
}

func main() {
	// true case:
	// crypt := []string{"SEND", "MORE", "MONEY"}
	// solution := [][]string{{"O", "0"},
	// 	{"M", "1"},
	// 	{"Y", "2"},
	// 	{"E", "5"},
	// 	{"N", "6"},
	// 	{"D", "7"},
	// 	{"R", "8"},
	// 	{"S", "9"}}

	// false case, because of leading zeros:
	// crypt := []string{"TEN", "TWO", "ONE"}
	// solution := [][]string{{"O", "1"},
	// 	{"T", "0"},
	// 	{"W", "9"},
	// 	{"E", "5"},
	// 	{"N", "4"}}

	// false case, zeroes
	// crypt := []string{"AA",
	// 	"AA",
	// 	"AA"}
	// solution := [][]string{{"A", "0"}}

	// false case, zeroes
	// crypt := []string{"A",
	// 	"A",
	// 	"A"}
	// solution := [][]string{{"A", "0"}}

	// false case, zeroes
	crypt := []string{"AA",
		"BB",
		"AA"}
	solution := [][]string{{"A", "1"}, {"B", "0"}}

	fmt.Println(isCryptSolution(crypt, solution))
}
