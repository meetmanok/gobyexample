package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("-------------Default value------------")
	var name string
	fmt.Println("name "+ name)
	var age int
	fmt.Println("age "+ strconv.Itoa(age))
	var balance float64
	fmt.Println("balance "+ strconv.FormatFloat(balance, 'f', 3, 32))
	var isEmployed bool
	fmt.Println("isEmployed "+ strconv.FormatBool(isEmployed))


	fmt.Println("-------------Value assignment------------")
	var fName string = "Manoj"
	lName := "Krishnan"
	fmt.Println(fName +" "+lName)

	var num1 int = 10
	num2 := 20
	fmt.Println(num1 + num2)

	var isTrue bool = true
	fmt.Println(isTrue && false)
}
