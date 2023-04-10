package conditions

import (
	"fmt"
	"math/rand"
	"time"
)

func Condition() {

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
