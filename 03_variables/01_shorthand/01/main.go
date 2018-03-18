package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%+v \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%T \n", d)
}
