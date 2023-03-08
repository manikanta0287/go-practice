package basics

import (
	"fmt"
	"reflect"
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

//String to boolean Integer or Float type

func MultilpleConversion() {

	str := "GeeksforGeeks"

	fmt.Println("Before :", reflect.TypeOf(str))
	fmt.Println("String value is: ", str)
	b, _ := strconv.ParseBool(str)
	fmt.Println("After :", reflect.TypeOf(b))
	fmt.Println("Boolean value is: ", b, "\n")

	str1 := "t"

	fmt.Println("Before :", reflect.TypeOf(str1))
	fmt.Println("String value is: ", str1)
	b1, _ := strconv.ParseBool(str1)
	fmt.Println("After :", reflect.TypeOf(b1))
	fmt.Println("Boolean value is: ", b1, "\n")

	str2 := "0"

	fmt.Println("Before :", reflect.TypeOf(str2))
	fmt.Println("String value is: ", str2)
	b2, _ := strconv.ParseBool(str2)
	fmt.Println("After :", reflect.TypeOf(b2))
	fmt.Println("Boolean value is: ", b2, "\n")

}

// convert String into Integer Data type

func StoInteger() {

	str := "GeeksforGeeks"

	fmt.Println("Before :", reflect.TypeOf(str))
	fmt.Println("String value is: ", str)
	i, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println("After :", reflect.TypeOf(i))
	fmt.Println("Integer value is: ", i, "\n")

	str1 := "100"

	fmt.Println("Before :", reflect.TypeOf(str1))
	fmt.Println("String value is: ", str1)
	i1, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Println("After :", reflect.TypeOf(i1))
	fmt.Println("Integer value is: ", i1, "\n")

}
