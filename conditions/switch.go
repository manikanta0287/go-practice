package conditions

import (
	"fmt"
)

func Switch() {
	days := 3

	switch days {
	case 1:
		fmt.Println("Enter the")
	case 2:
		fmt.Println("Week is a Tuesday")
	case 3:
		fmt.Println("Week is a Wednesday")
	default:
		fmt.Println("Jagan wants holiday")
	}
}

func Weekdays() {
	var letter string
	fmt.Print("Enter the letter : ")
	fmt.Scanf("%s", &letter)
	fmt.Println("letter", letter)

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend Go and Enjoy Your weekend")
	// default:
	// 	fmt.Println("come to office and do your work")
	// }

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("Not vowel")
	}

}

func AnyValue(value map[int]string) map[int]string {
	// value = make(map[int]string)

	fmt.Println("value", value)
	return value
}
