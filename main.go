package main

import (
	a "example/hello/anagram"
	binarytree "example/hello/binaryTree"
	p "example/hello/pattern"
	"fmt"
	"strconv"
)

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
	fmt.Scanln(&n)
	fmt.Println("Factorial of ", n, " is ", fact(n))
}

func merge(left []int, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		merged = append(merged, left[i])
	}
	for ; j < len(right); j++ {
		merged = append(merged, right[j])
	}

	return merged
}

func devide(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	middle := len(arr) / 2
	fmt.Println("Left", arr[:middle])
	left := devide(arr[:middle])

	fmt.Println("Right", arr[middle:])
	right := devide(arr[middle:])

	return merge(left, right)
}

func mergeSort() {
	var n int
	mergeSlice := []int{}
	fmt.Println("Enter the number of elements to be sorted:")
	fmt.Scan(&n)
	fmt.Println("Enter the numbers")

	var m int
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		mergeSlice = append(mergeSlice, m)
	}

	fmt.Println("Sorted: ", devide(mergeSlice))
}

func drawTriangle() {
	var n int

	fmt.Println("Enter n:")
	fmt.Scan(&n)
	fmt.Println("***********Pattern using maps**********")

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := i; k > 0; k-- {
			fmt.Print(strconv.Itoa(k))
		}
		if i != 1 {
			for l := 2; l <= i; l++ {
				fmt.Print(strconv.Itoa(l))
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	var i int
	for {
		fmt.Println("\n*****Choose the function you want to perform:")
		fmt.Println("1: Draw Pattern")
		fmt.Println("2: Factorial of n")
		fmt.Println("3: Merge sort")
		fmt.Println("4: Draw triangle")
		fmt.Println("5: Binary Tree")
		fmt.Println("6: Anagram")
		fmt.Scan(&i)

		switch i {
		case 1:
			p.DrawPatternUsingMaps()
		case 2:
			factorial()
		case 3:
			mergeSort()
		case 4:
			drawTriangle()
		case 5:
			binarytree.ImplementBinaryTree()
		case 6:
			a.Anagram()
		default:
			fmt.Println("Enter the correct option number")
		}

	}
}
