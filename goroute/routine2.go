package goroute

import (
	"fmt"
	"time"
)

func One() {
	fmt.Println("one")

	fmt.Println("two")

	fmt.Println("three")
	time.Sleep(10 * time.Second)
}
