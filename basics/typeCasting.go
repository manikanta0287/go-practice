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

func TypeCon() {

	// taking the required
	// data into variables
	var totalsum int = 846
	var number int = 19
	var avg float32

	// explicit type conversion
	avg = float32(totalsum) / float32(number)

	// Displaying the result
	fmt.Printf("Average = %f\n", avg)
}
