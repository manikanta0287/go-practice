package conditions

import (
	"fmt"
	"math/rand"
	"time"
)

func Condition() {
	// In this example, we generate random values between -5 and 4. With the help of the conditionals, we print a message for all three options.
	rand.Seed(time.Now().UnixNano())

	num := -5 + rand.Intn(10)

	if num > 0 {

		fmt.Println("The number is positive")
	} else if num == 0 {

		fmt.Println("The number is zero")
	} else {

		fmt.Println("The number is negative")
	}
}
