package logics

import "fmt"

/* func Stars() int {

	var rows, columns int

	fmt.Print("nEnter the number of rows : ")

	fmt.Scanf("%d", &rows)

	fmt.Printf("nEnter the number of columns : ")

	fmt.Scanf("%d", &columns)

	fmt.Printf("")

	Rect(rows, columns)

	return 0

}

func Rect(x int, y int) {

	for i := 1; i <= x; i++ {

		for j := 1; j <= y; j++ {

			fmt.Printf(" * ")

		}

		fmt.Printf("\n")

	}

} */

//Printing Stars using single functions

func PrintStarts() {
	var x, y int

	fmt.Print("Enter the number of rows: ")
	fmt.Scanf("%d", &x)
	fmt.Print("Enter the number of columns: ")
	fmt.Scanf("%d", &y)
	for i := 1; i <= x; i++ {

		for j := 1; j <= y; j++ {

			fmt.Printf(" * ")

		}

		fmt.Printf("\n")

	}
}