package basics

func Abc(x int) int {
	return x + 1
}

func Multireturn(x int) (int, int){
	return x, x+1
}


func Add(x int, y int, value string) (int, string) {
	return x + y, value
}
