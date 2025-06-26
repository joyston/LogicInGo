package pattern

import (
	"fmt"
	"strconv"
)

func DrawPattern() {
	var n int

	fmt.Println("Enter n:")
	fmt.Scan(&n)
	fmt.Println("***********Pattern**********")

	if n == 1 {
		fmt.Println("1")
	} else {
		rows := (2 * n)
		for i := 1; i <= rows-1; i++ {
			for j := 1; j <= n; j++ {
				if i < n {
					fmt.Print(" ")
				} else {
					fmt.Println(i)
				}
			}

		}
	}
}

func DrawPatternUsingMaps() {
	var n int

	fmt.Println("Enter n:")
	fmt.Scan(&n)
	fmt.Println("***********Pattern using maps**********")

	if n == 1 {
		fmt.Println("1")
	} else {
		rows := (2 * n)

		var pattern = make(map[int]string)
		for i := 1; i <= rows; i++ {
			pattern[i] = ""

			//add space
			for k := n - i; k > 0; k-- {
				pattern[i] = pattern[i] + " "
			}

			j := i
			for len(pattern) <= n {
				if j == 1 {
					pattern[i] = pattern[i] + strconv.Itoa(j)
					j++
				}

				if j > i {
					break
				} else if j > 1 {
					pattern[i] = pattern[i] + strconv.Itoa(j)
					j--
				} else {
					pattern[i] = pattern[i] + strconv.Itoa(j)
					j++
				}

			}

		}
		for _, v := range pattern {
			fmt.Println(v)
		}
	}
}
