package main

//The standard library’s strings package provides many useful string-related functions.
//Here are some examples to give you a sense of the package.

import (
	"fmt"
	s "strings"
)

//We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = fmt.Println

func main(){
	//Here’s a sample of the functions available in strings. Since these are functions from the package,
	//not methods on the string object itself, we need pass the string in question as the first argument to the function.
	//You can find more functions in the strings package docs.
	p("Contains:", s.Contains("Manoj", "a"))


	p("HasPrefix:", s.HasPrefix("K.Manoj", "K."))
	p("HasSuffix:", s.HasSuffix("Manoj.K", ".K"))
	p("HasSuffix:", s.HasSuffix("Manoj.K", ".M"))

	p("String Repeat:", s.Repeat("manoj", 3))

	p("Split:", s.Split("a-b-c-d-e", "-"))
	p("Join:", s.Join([]string{"a", "b", "c", "d", "e"}, "_"))

	//Replace
	p("Replace Unlimited:", s.Replace("foo", "o", "O", -1))
	p("Replace only 1st occurrence:", s.Replace("foo", "o", "O", 1))

	//Compare
	p("Compare:", s.Compare("Manoj", "Manoj"))
	p("Compare:", s.Compare("Manoj", "Manoj.K"))
	p("Compare:", s.Compare("Manoj.K", "Manoj"))

	p("ToLower:",s.ToLower("MANOJ.K"))
	p("ToLower:",s.ToUpper("manoj.k"))

	//Not part of strings, but worth mentioning here,
	//are the mechanisms for getting the length of a string in bytes and getting a byte by index.
	p("Length of String:",len("manoj.k"))
	p("Char:","manoj.k"[0:2])
}
