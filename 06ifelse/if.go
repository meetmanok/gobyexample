package main

/*
Branching with if and else in Go is straight-forward.
 */
import "fmt"

func main() {
	//A statement can precede conditionals; any variables declared in this statement are available in all branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//You can have an if statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//Normal if else
	num := 9
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//Note: that you don’t need parentheses around conditions in Go, but that the braces are required.
	//Note: There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
}
