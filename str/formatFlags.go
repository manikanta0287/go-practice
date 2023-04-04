package str

import "fmt"

func FormatFlags() {

	fmt.Printf("%+d\n", 1691)

	fmt.Println("---------------------")

	fmt.Printf("%#x\n", 1691)
	fmt.Printf("%#X\n", 1691)
	fmt.Printf("%#b\n", 1691)

	fmt.Println("---------------------")

	fmt.Printf("%10d\n", 1691)
	fmt.Printf("%-10d\n", 1691)
	fmt.Printf("%010d\n", 1691)
}
