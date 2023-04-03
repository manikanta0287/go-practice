package str

import (
	"fmt"
)

func BasicString() {

	name := "Jane"
	age := 17

	fmt.Printf("%s is %d years old\n", name, age)

	res := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(res)
}
