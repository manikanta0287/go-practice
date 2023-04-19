// You can edit this code!
// Click here and start typing.
package maps

import "fmt"

func Mapping() {
	m := make(map[int]string)
	push(1, "Ganesh", m)
	push(2, "Ganesh1", m)
	push(3, "Ganesh2", m)
	push(4, "Ganesh3", m)
	push(5, "Ganesh4", m)
	read(m)
	pop(2, m)
	read(m)
	fmt.Println("Hello!!!")
}

func push(key int, value string, m map[int]string) {
	m[key] = value
}
func pop(key int, m map[int]string) {
	delete(m, key)
}

func read(m map[int]string) {
	fmt.Println(m)
}

//-----------------------------------------------------------

func TestMap() {
	v := make(map[int]string)

	v[1] = "abc"
	fmt.Println(v[1])

	v[4] = "Hello"
	v[16] = "World"
	fmt.Println("v[2] :", v[2], v[16])

	fmt.Println("length : ", len(v))

	//------------Direct Assigning---------------------------

	a := map[int]string{22: "Hello", 23: "World"}

	//Concatinating here----
	fmt.Println("a>>>>>>>:", a[22]+a[23])

	//------------map Interface-------------------------------

	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	fmt.Println("foods>>>:", foods)

}
