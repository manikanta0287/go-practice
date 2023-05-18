package conditions

import "fmt"

func Loop2() {
	// Ranging
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0

	for _, num := range nums {
		sum += num
		fmt.Println("num >>>>", num)
	}

	fmt.Println(sum)

	for _, name := range nums {

		fmt.Println("Hello", name)
	}
}
