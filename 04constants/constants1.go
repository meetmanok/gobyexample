package main

import (
	"fmt"
	"strconv"
)

const (
	name string = "manoj"
	age = 28
)

func main(){
	fmt.Println("name = ", name)
	fmt.Printf("%T\n", name)
	fmt.Println("age = ", strconv.Itoa(age))
	fmt.Printf("%T\n", age)

	const name = "Logu"
	fmt.Println("name = ", name)
	fmt.Printf("%T\n", name)
}
