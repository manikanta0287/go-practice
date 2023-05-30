package goroute

import "fmt"

func SendData(ch chan<- int) {
	// ch := make(chan int)
	// ch <- 20
	//x := <-ch
	// ch <- 1
	// ch <- 1
	// ch <- 1
	fmt.Println("ch>>>>>>>>>>", ch)
	// close(ch)
	fmt.Println("ch>>>>>>>>>>", ch)

}
