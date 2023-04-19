package conditions

import (
	"fmt"
	"time"
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
	// var day int
	// fmt.Println("Enter the day value: ")
	// fmt.Scanf("%d", &day)
	// fmt.Println("day", day)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend Go and Enjoy Your weekend")
	default:
		fmt.Println("come to office and do your work")
	}

}

func AnyValue(value map[int]string) map[int]string {
	// value = make(map[int]string)

	fmt.Println("value", value)
	return value
}
