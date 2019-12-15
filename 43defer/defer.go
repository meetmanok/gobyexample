package main

/*
Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
defer is often used where e.g. ensure and finally would be used in other languages.
 */

import (
	"fmt"
	"os"
)

//Suppose we wanted to create a file, write to it, and then close when we’re done. Here’s how we could do that with defer.
func main(){
	//Immediately after getting a file object with createFile, we defer the closing of that file with closeFile.
	//This will be executed at the end of the enclosing function (main), after writeFile has finished.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(file string) *os.File{
	fmt.Println("Creating File...")
	f, err := os.Create(file)
	if err != nil {
		fmt.Println("Failed to create file...")
		os.Exit(1)
	}
	return f
}

func closeFile(file *os.File) {
	fmt.Println("Closing File...")
	err := file.Close()
	//It’s important to check for errors when closing a file, even in a deferred function.
	if err != nil {
		fmt.Println("Failed to close file...")
		os.Exit(1)
	}
}

func writeFile(file *os.File) {
	fmt.Println("Writing File...")
	fmt.Fprintln(file, "data")
}

//Running the program confirms that the file is closed after being written.
