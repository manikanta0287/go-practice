package logics

import "fmt"

// C program to print half pyramid pattern using stars
func HalfPyramid() {
	var n int
	fmt.Print("Enter the n: ")
	fmt.Scanf("%d", &n)
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" * ")
		}
		fmt.Print("\n")
	}
}

//Second Method

/* func HalfPyramid() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
	fmt.Println("END")
} */

//labels

func Label() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println("category=", i, "product=", j)
		}
	}
	fmt.Println("END")
}

//Label using Break condition

func LableWithBreak() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println("i=", i, "j=", j)
			if i == j {
				break
			}
		}
	}
}
