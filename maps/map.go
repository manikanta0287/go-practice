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
