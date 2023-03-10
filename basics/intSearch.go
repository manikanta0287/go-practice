package basics

import (
	"fmt"
	"sort"
)

func IntSearch() {
	s1 := []int{23, 54, 21, 65, 87, 13}

	s2 := []int{400, 200, 123, 500, 654}

	fmt.Println(s1, s2)

	sort.Ints(s1)
	if sort.IntsAreSorted(s1) {
		fmt.Println("Inst are sorted")
	} else {
		fmt.Println("Error in sorting")
	}

	fmt.Println(s1)
}
