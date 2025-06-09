package main

import "fmt"

func drawPattern() {
	var n int

	fmt.Println("Enter n:")
	fmt.Scan(&n)
	fmt.Println("***********Pattern**********")

	if n == 1 {
		fmt.Println("1")
	} else {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if j < n {
					fmt.Print(" ")
				} else {
					fmt.Println(i)
				}
			}

		}
	}
}

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func factorial() {
	var n int

	fmt.Println("Enter n to find it's factorial:")
	fmt.Scan(&n)
	fmt.Println("Factorial of ", n, " is ", fact(n))
}

func main() {
	var i int
	for {
		fmt.Println("*****Choose the function you want to perform:")
		fmt.Println("1: Draw Pattern")
		fmt.Println("2: Factorial of n")
		fmt.Scan(&i)

		switch i {
		case 1:
			drawPattern()
		case 2:
			factorial()
		default:
			fmt.Println("Enter the correct option number")
		}

	}
}
