package main

import (
	"fmt"
)

// Spent 22 minutes
func firstNotRepeatingCharacter(s string) string {

	res := "_"
	type S struct {
		c int
		i int
	}

	h := make([]S, 255)

	for i := 0; i < len(s); i++ {
		h[s[i]].c++
		h[s[i]].i = i
		// fmt.Printf("%c - [%#v]\n", s[i], h[s[i]])
	}

	m := -1
	for i, c := range h {
		if c.c == 1 {
			if m > c.i || m == -1 {
				m = c.i
				res = fmt.Sprintf("%c", i)
				// fmt.Printf(">>> %d \t %#v \t %c\n", i, c, i)
			}
		}
	}

	// fmt.Printf("%#v\n", h)

	return res
}

func main() {

	fmt.Println(firstNotRepeatingCharacter("abacabad"))
	fmt.Println(firstNotRepeatingCharacter("abacabaabacaba"))
	fmt.Println(firstNotRepeatingCharacter("ngrhhqbhnsipkcoqjyviikvxbxyphsnjpdxkhtadltsuxbfbrkof"))
}
