package goroute

import (
	"fmt"
	"sync"
)

func PrintWaitGroup(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	
	words := []string{
		"alpha",
		"beta",
		"gama",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
	}
	wg.Add(8)
	
	for i, v := range words {
		go PrintWaitGroup(fmt.Sprintf("%d, %s", i, v), &wg)
	}
	wg.Wait()
}

//Please convert this according to architecture