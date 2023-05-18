package conditions

import "fmt"

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
	for i := 10; i < 20; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println(i * j)
			// if i == j {
			// 	break
			// }
		}
	}
}
