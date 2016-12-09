package main

import "fmt"

func main() {
	var a int
	a = 10
	fmt.Println(a)
	duplicar(&a)
	fmt.Println(a)
}

func duplicar(x *int) {
	//fmt.Println(x)
	*x = *x * 2
}
