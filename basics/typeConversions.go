package basics

import (
	"fmt"
	"strconv"
)

func Conversion() {
	// fmt.Println("Hai@@@@")

	a := true
	//converting from Boolean to string
	s := strconv.FormatBool(a)
	fmt.Printf("%T\n", a)
	fmt.Println(s)

	//2. fmt.Sprintf() Method
	b := false

	p := fmt.Sprintf("value is :%v\n and type is:%T\n", b, b)
	fmt.Println(p)

}
