package basics

import (
	"fmt"
	"sort"
)

func IntSearch() {
	s1 := []int{23, 54, 21, 65, 87, 13}

	s2 := []string{"a", "z", "t", "w", "h"}

	fmt.Println(s1, s2)

	sort.Ints(s1)
	if sort.IntsAreSorted(s1) {
		fmt.Println("Inst are sorted")
	} else {
		fmt.Println("Error in sorting")
	}

	sort.Strings(s2)

	if sort.StringsAreSorted(s2) {
		fmt.Println("Strings are sorted")
	} else {
		fmt.Println("Error in sorting")
	}
	fmt.Println(s1, s2)

	// reverse := sort.Reverse(s1)
	// fmt.Println(reverse)
}
