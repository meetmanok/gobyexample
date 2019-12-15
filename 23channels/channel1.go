package main

import (
	"fmt"
)

func main(){
	c := make(chan int)

	go foo(c)

	bar(c)

	println("About to exit...")
}

func foo(c chan <- int){
	c <- 20
}

func bar(c <- chan int){
	fmt.Println(<-c)
}
