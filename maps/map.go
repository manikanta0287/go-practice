// You can edit this code!
// Click here and start typing.
package maps

import (
	"encoding/json"
	"fmt"
	"log"
)

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

	person := map[string]interface{}{
		"name": "manikanta",
		"in-details": struct {
			age         int
			location    string
			designation string
		}{33, "Hyd", "SE"},
		"technology": "Go",
	}
	fmt.Println("person map is :", person)

}

//Declaring a map

func DeclareMap() {
	var menu map[string]float64

	menu = map[string]float64{
		"eggs":  1.75,
		"bacon": 3.22,
		"saus":  1.89,
	}
	fmt.Println("menu is :", menu)

	type menuItems struct {
		price float64
	}

	var menu2 = map[string]menuItems{
		"beans": menuItems{
			price: 0.49,
		},
	}
	fmt.Println("menu2 is:", menu2)

	// beans := menu["beans"]
	// beans.price = 0.25
	// menu["beans"] = beans

	/* var menu3 = map[string]bool{
		"chicken": true,
		"steak":   true,
	}
	if _, ok := menu3[dish]; ok {
		fmt.Sprintf("Yes , %s is in menu2", dish)
	}
	fmt.Sprintln("Sorry , we're all out of %s", dish)
	fmt.Println(menu3) */

	//Iterating the map
	for item, price := range menu {
		fmt.Println(item, price)
	}

	//Copying a map
	m := map[string]bool{
		"abc": true,
		"bcd": false,
		"xyz": true,
	}
	fmt.Println(m)

	c := map[string]bool{}
	for k, v := range m {
		c[k] = v
	}

	fmt.Println(c)

	//Convert map into json data
	Map := map[string]bool{
		"abc": true,
		"bcd": false,
		"xyz": true,
	}
	data, err := json.Marshal(Map)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
