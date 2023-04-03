package str

import "fmt"

func Index() {

	n1 := 2
	n2 := 3
	n3 := 4

	// Observe the indexing
	res := fmt.Sprintf("There are %d oranges %d apples %d plums", n1, n2, n3)
	fmt.Println(res)

	res2 := fmt.Sprintf("There are %[2]d oranges %d apples %[1]d plums", n1, n2, n3)
	fmt.Println(res2)

	res3 := fmt.Sprintf("There are %[1]d oranges %[2]d apples %d plums", n1, n2, n3)
	fmt.Println(res3)
}
