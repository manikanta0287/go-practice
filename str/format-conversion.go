package str

import "fmt"

func FormatConversion() {

	fmt.Printf("%d\n", 1671)
	fmt.Printf("%o\n", 1671)
	fmt.Printf("%x\n", 1671)
	fmt.Printf("%X\n", 1671)
	fmt.Printf("%#b\n", 1671)
	fmt.Printf("%f\n", 1671.678)
	fmt.Printf("%F\n", 1671.678)
	fmt.Printf("%e\n", 1671.678)
	fmt.Printf("%E\n", 1671.678)
	fmt.Printf("%g\n", 1671.678)
	fmt.Printf("%G\n", 1671.678)
	fmt.Printf("%s\n", "Zetcode")
	fmt.Printf("%c %c %c %c %c %c %c\n", 'Z', 'e', 't', 'C', 'o', 'd', 'e')
	fmt.Printf("%p\n", []int{1, 2, 3})
	fmt.Printf("%d%%\n", 1671)
	fmt.Printf("%t\n", 3 > 5)
	fmt.Printf("%t\n", 5 > 3)
}
