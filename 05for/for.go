package main

import (
	"fmt"
	"strconv"
)

/*
for is Go’s only looping construct. Here are three basic types of for loops.
 */
func main(){
	//A classic initial/condition/after for loop.
	//Normal loop
	for i := 0; i <= 10; i++{
		fmt.Printf("%v\t%#U\n", i, i)
	}

	//Loop through Array
	a := []int{1,2,3,4,5,6,7,8,9}
	for i, v := range a {
		fmt.Println("index="+strconv.Itoa(i)+", value="+strconv.Itoa(v))
	}

	//Loop through Map
	ages := map[string]int{
		"manoj": 29,
		"logu": 25,
	}
	for k, v := range ages{
		fmt.Println("key=" + k + ", value=" + strconv.Itoa(v))
	}

	//The most basic type, with a single condition.
	//Like while loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	//Infinite loop
	for {
		fmt.Println("Infinite loop")
		break
	}
	fmt.Println("After break...")

	//return will also exit from infinite loop but next statement wont execute in this scenario
	for {
		fmt.Println("Infinite loop")
		return
	}
	fmt.Println("After return, this wont be printed...")

	//We’ll see some other for forms later when we look at range statements, channels, and other data structures.
}