package main

import (
	"fmt"
	// "math/rand"
)

func firstDuplicateSlow(a []int) int {
	// A: 2 3 3 1 5 2
	//    0 1 2 3 4 5
	// i: x
	// j:           x

	ret := -1
	data := -1
	counter := 0
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				fmt.Printf("%d %d\t ret: %d\t a[i]: %d\t a[j]: %d\n", i, j, ret, a[i], a[j])
				counter++
				if j < ret || ret == -1 {
					ret = j
					data = a[i]
				}
			}
		}
	}
	fmt.Printf("Array: [%d], Counter: [%d]\n", len(a), counter)

	return data
}

func firstDuplicate(a []int) int {
	//    0 1 2 3 4 5
	// a: 2 3 3 1 5 2
	// a: 1 1 2 2 1

	//    0 1 2 3 4 5
	// h: 0 1 2 2 0 1
	// h:   3 5 2   4

	//    0 1 2 3 4 5
	// h:   2 2
	// h:

	type t struct {
		v int
		r int
		c int
	}
	h := make([]t, 100000)
	for i, _ := range a {
		if h[a[i]].c < 2 {
			h[a[i]].v = a[i]
			h[a[i]].r = i
			h[a[i]].c++
		}
	}

	fmt.Println(a)
	left := -1
	value := -1
	for _, item := range h {
		if item.c > 1 {
			if item.r < left || left == -1 {
				left = item.r
				value = item.v
			}
		}
		fmt.Printf("v:%d; r:%d; c:%d\n", item.v, item.r, item.c)
	}

	return value
}

func main() {
	// ar := []int{2, 3, 3, 1, 5, 2}
	ar := []int{1, 1, 2, 2, 1}
	// ar := make([]int, 1000)
	// for i, _ := range ar {
	// 	ar[i] = rand.Int() % 100
	// }
	fmt.Println(firstDuplicate(ar))
}
