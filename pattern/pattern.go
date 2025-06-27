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
		// rows := (2 * n)

		var pattern = make(map[int]string)
		for i := 1; i <= n; i++ {

			pattern[i] = ""

			//add space
			for j := 0; j < n-i; j++ {
				pattern[i] = pattern[i] + " "
			}

			for k := i; k > 1; k-- {
				pattern[i] = pattern[i] + strconv.Itoa(k)
			}

			for m := 1; m <= i; m++ {
				pattern[i] = pattern[i] + strconv.Itoa(m)
			}
			fmt.Println(pattern[i])
		}

		for p := len(pattern) - 1; p > 0; p-- {
			fmt.Println(pattern[p])
		}
	}
}
