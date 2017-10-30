package main

import (
	"fmt"
)

func rotateImage(a [][]int) [][]int {
	J := len(a)
	H := make([][]int, J)

	for i, _ := range H {
		H[i] = make([]int, len(a[0]))
	}

	for i, r := range a {
		for j, _ := range r {
			// fmt.Printf("%d; ", a[J-j-1][i])
			H[i][j] = a[J-j-1][i]
		}
		// fmt.Println()
	}

	return H
}

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	fmt.Println(a)

	r := rotateImage(a)

	fmt.Println(r)
}
