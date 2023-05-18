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

// Using fall through we can execute false value also
func FallThrough() {

	switch num := 25; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
	}

}

// Using Break in switch
func BreakSwitch() {
	switch num := -5; {
	case num < 50:
		if num < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}
