package main

import (
	"fmt"
)

func hashCount(a []string) bool {
	h := make([]int, 255)

	for i, _ := range a {
		if a[i] != "." {
			fmt.Printf("[%c:%d] ", a[i][0], h[a[i][0]])
			if h[a[i][0]] > 0 {
				return false
			}
			h[a[i][0]]++
		}
	}
	fmt.Println()

	return true
}

func sudoku2(grid [][]string) bool {

	// Addressing: grid[column][row]

	// Horizontal:
	fmt.Println("Horizontal")
	for _, h := range grid {
		b := hashCount(h)
		if b == false {
			fmt.Println(h)
			return false
		}
	}

	// Vertical:
	fmt.Println("Vertical")
	columnAr := make([]string, len(grid))
	for c := 0; c < len(grid[0]); c++ {

		for r := 0; r < len(grid); r++ {
			columnAr[r] = grid[r][c]
		}

		b := hashCount(columnAr)

		if b == false {
			fmt.Println(columnAr)
			return false
		}
	}

	// Vertical:
	fmt.Println("3x3")
	flat3x3 := make([]string, 9)
	i := 0
	for r := 0; r < len(grid)-2; r += 3 {
		for c := 0; c < len(grid[0])-2; c += 3 {
			i = 0
			for r0 := 0; r0 < 3; r0++ {
				for c0 := 0; c0 < 3; c0++ {
					flat3x3[i] = grid[r+r0][c+c0]
					i++
				}
			}

			b := hashCount(flat3x3)
			if b == false {
				fmt.Println(flat3x3)
				return false
			}

		}
	}

	return true
}

func main() {

	// g := [][]string{
	// 	{".", ".", ".", "1", "4", ".", ".", "2", "."},
	// 	{".", ".", "6", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", ".", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", "1", ".", ".", ".", ".", ".", "."},
	// 	{".", "6", "7", ".", ".", ".", ".", ".", "9"},
	// 	{".", ".", ".", ".", ".", ".", "8", "1", "."},
	// 	{".", "3", ".", ".", ".", ".", ".", ".", "6"},
	// 	{".", ".", ".", ".", ".", "7", ".", ".", "."},
	// 	{".", ".", ".", "5", ".", ".", ".", "7", "."}}

	// g2 := [][]string{
	// 	{".", "9", ".", ".", "4", ".", ".", ".", "."},
	// 	{"1", ".", ".", ".", ".", ".", "6", ".", "."},
	// 	{".", ".", "3", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", ".", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", ".", "7", ".", ".", ".", ".", "."},
	// 	{"3", ".", ".", ".", "5", ".", ".", ".", "."},
	// 	{".", ".", "7", ".", ".", "4", ".", ".", "."},
	// 	{".", ".", ".", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", ".", ".", "7", ".", ".", ".", "."}}

	g3 := [][]string{
		{".", ".", ".", ".", ".", ".", "5", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"9", "3", ".", ".", "2", ".", "4", ".", "."},
		{".", ".", "7", ".", ".", ".", "3", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", "3", "4", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "3", ".", ".", "."},
		{".", ".", ".", ".", ".", "5", "2", ".", "."}}

	// fmt.Println(sudoku2(g))
	// fmt.Println(sudoku2(g2))
	fmt.Println(sudoku2(g3))
}
