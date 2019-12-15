package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("Hello, I am Gopher!!!")
	}

	v := func() int {
		return 10
	}()

	fmt.Println(v)
}
