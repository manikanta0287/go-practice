package str

import "fmt"

type User struct {
	name       string
	occupation string
}

func GeneralVerb() {

	msg := "and old falcon"
	n := 16
	w := 12.45
	r := true
	u := User{"John Doe", "gardener"}
	vals := []int{1, 2, 3, 4, 5}
	ctrs := map[string]string{
		"sk": "Slovakia",
		"ru": "Russia",
		"de": "Germany",
		"no": "Norway",
	}

	fmt.Printf("%v %v %v %v %v\n  %v %v\n", msg, n, w, u, r, vals, ctrs)
	fmt.Printf("%v %+v\n", u, u)

	fmt.Println("--------------------")

	fmt.Printf("%#v %#v %#v %#v %#v\n  %#v %#v\n", msg, n, w, u, r, vals, ctrs)
	fmt.Printf("%T %T %T %T %T %T %T\n", msg, n, w, u, r, vals, ctrs)

	fmt.Println("--------------------")

	fmt.Printf("The prices dropped by 12%%\n")
}
