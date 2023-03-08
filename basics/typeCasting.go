package basics

import "fmt"

func TypeCasting() {

	//Check the data type of a and b, so the data conversion gives compile error
	// var a int64 = 200
	// var b int = a
	// fmt.Println("test", a, b)

	var a int = 200
	var b int = a
	fmt.Println("test", a, b)
}
