package main

import (
	"bytes"
	"fmt"
	"regexp"
)

/*
Go offers built-in support for regular expressions. Here are some examples of common regexp-related tasks in Go.
 */

func main(){
	//This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	// o/p: true

	//Above we used a string pattern directly, but for other regexp tasks you’ll need to Compile an optimized Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	//Many methods are available on these structs. Here’s a match test like we saw earlier.
	fmt.Println(r.MatchString("peach"))
	// o/p: true

	//This finds the match for the regexp.
	fmt.Println(r.FindString("peach punch"))
	// o/p: peach

	//This also finds the first match but returns the start and end indexes for the match instead of the matching text.
	fmt.Println(r.FindStringIndex("peach punch"))
	// o/p: [0 5]

	//The Submatch variants include information about both the whole-pattern matches and the submatches within those matches.
	//For example this will return information for both p([a-z]+)ch and ([a-z]+).
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// o/p: [peach ea]

	//Similarly this will return information about the indexes of matches and submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// o/p: [0 5 1 3]

	//The All variants of these functions apply to all matches in the input, not just the first. For example to find all matches for a regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// o/p: [peach punch pinch]

	//These All variants are available for the other functions we saw above as well.
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// o/p: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	fmt.Println(r.FindAllStringIndex("peach punch pinch", -1))
	// o/p: [[0 5] [6 11] [12 16]]

	//Providing a non-negative integer as the second argument to these functions will limit the number of matches.
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// o/p: [peach punch]

	//Our examples above had string arguments and used names like MatchString.
	//We can also provide []byte arguments and drop String from the function name.
	fmt.Println(r.Match([]byte("peach")))
	// o/p: true

	//When creating constants with regular expressions you can use the MustCompile variation of Compile.
	//A plain Compile won’t work for constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	// o/p: p([a-z]+)ch

	//The regexp package can also be used to replace subsets of strings with other values.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// o/p: a <fruit>

	//The Func variant allows you to transform matched text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
	// o/p: a PEACH
}

//For a complete reference on Go regular expressions check the regexp package docs.