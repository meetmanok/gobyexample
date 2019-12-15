package main

/*
Functions are central in Go. We’ll learn about functions with a few different examples.
 */

import "fmt"


//Here’s a function that takes two ints and returns their sum as an int.
func plus(num1, num2 int) int {
	return num1 + num2
}

//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
func plus1(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//When you have multiple consecutive parameters of the same type,
//you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plus2(num1, num2 int) {
	fmt.Println(num1 + num2)
}

func main(){
	//Call a function just as you’d expect, with name(args).
	res := plus(10, 20)
	fmt.Println(res)

	res = plus1(10, 20, 30)
	fmt.Println(res)

	plus2(10, 20)
}

//There are several other features to Go functions. One is multiple return values, which we’ll look at next.