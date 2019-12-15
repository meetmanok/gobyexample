package main

/*
Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
 */

import (
	"fmt"
)

func main() {
	person := map[string]int{
		"manoj": 29,
		"sindhu": 28,
		"Logu": 26,
	}
	fmt.Println(person)

	//Set key/value pairs using typical name[key] = val syntax.
	//Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	//Add new element
	person["gobi"] = 23
	person["ashok"] = 27
	fmt.Println(person)

	//Check element is there in map
	v, ok := person["gobi"]
	fmt.Printf("%v %v\n",v, ok)
	v, ok = person["modi"]
	fmt.Printf("%v %v\n",v, ok)

	//Delete element in map
	delete(person, "gobi")
	fmt.Println(person)

	//To create an empty map, use the builtin make: make(map[key-type]val-type).
	//Using make
	person = make(map[string]int)
	person["manoj"] = 29
	person["sindhu"] = 28
	person["logu"] = 26
	fmt.Println(person)
}
