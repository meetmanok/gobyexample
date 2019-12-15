package main

/*
We often need our programs to perform operations on collections of data,
like selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function.

In some languages it’s idiomatic to use generic data structures and algorithms. Go does not support generics;
in Go it’s common to provide collection functions if and when they are specifically needed for your program and data types.

Here are some example collection functions for slices of strings. You can use these examples to build your own functions.
Note that in some cases it may be clearest to just inline the collection-manipulating code directly, instead of creating and calling a helper function.
 */
import (
	"fmt"
	"strings"
)

//Index returns the first index of the target string t, or -1 if no match is found.
func Index(arr []string, s string) int {
	for i, v := range arr{
		if v == s {
			return i
		}
	}
	return -1
}

//Include returns true if the target string s is in the slice.
func Include(arr []string, s string) bool {
	for _, v := range arr{
		if v == s {
			return true
		}
	}
	return false
}

//Any returns true if one of the strings in the slice satisfies the predicate f.
func Any(arr []string, f func(string) bool) bool {
	for _, v := range arr{
		if f(v) {
			return true
		}
	}
	return false
}

//All returns true if all of the strings in the slice satisfy the predicate f.
func All(arr []string, f func(string) bool) bool {
	for _, v := range arr{
		if !f(v) {
			return false
		}
	}
	return true
}

//Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(arr []string, f func(string) bool) []string {
	var newarr []string
	for _, v := range arr{
		if f(v){
			newarr = append(newarr, v)
		}
	}
	return newarr
}

//Map returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(arr []string, f func(string) string) []string {
	newarr := make([]string, len(arr))
	for i, v := range arr{
		newarr[i] = f(v)
	}
	return newarr
}

func main(){
	//Here we try out our various collection functions.
	strs := []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	//The above examples all used anonymous functions, but you can also use named functions of the correct type.
	fmt.Println(Map(strs, strings.ToUpper))
}