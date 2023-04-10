package conditions

import "fmt"

func Loop() {

	// num := 0

	// for i := 0; i < 5; i++ {
	// 	num = num + i
	// }
	// fmt.Println(num)

	sum := 0

	for i := 0; i < 5; i++ {
		sum += i
	}

	fmt.Println(sum)

}
