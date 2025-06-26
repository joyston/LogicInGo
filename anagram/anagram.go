package anagram

import (
	"fmt"
)

func Compare(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int)

	for _, Val := range s1 {
		_, ok := m[Val]
		if ok {
			m[Val]++
		} else {
			m[Val] = 1
		}
	}

	for _, Val := range s2 {
		_, ok := m[Val]
		if !ok {
			return false
		} else {
			if m[Val] == 1 {
				delete(m, Val)
			} else {
				m[Val]--
			}
		}
	}
	return true
}

func Anagram() {
	var s1, s2 string
	fmt.Println("Enter string S1:")
	fmt.Scan(&s1)
	fmt.Println("Enter string S2:")
	fmt.Scan(&s2)
	if Compare(s1, s2) {
		fmt.Println("TRUE: Strings S1 & S2 are Anagrams")
	} else {
		fmt.Println("FALSE: Strings S1 & S2 are NOT Anagrams")
	}
}
