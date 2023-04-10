package conditions

import "fmt"

func Switch() {
	days := 3

	switch days {
	case 1:
		fmt.Println("Week is a Monday")
	case 2:
		fmt.Println("Week is a Tuesday")
	case 3:
		fmt.Println("Week is a Wednesday")
	default:
		fmt.Println("Jagan wants holiday")
	}
}
