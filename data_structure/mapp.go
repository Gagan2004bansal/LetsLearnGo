package main

import (
	"fmt"
	"maps"
)

// maps -> hash, objects, dict
func mapp() {

	// create a map with string keys and string values
	m := make(map[string]string)

	// setting an element
	m["name"] = "Gagan"

	// getting an element
	fmt.Println("Name :", m["name"])

	// len of map
	fmt.Println(len(m))

	// creation of map without using make
	mpp := map[string]string{}
	mpp["name"] = "Gagan"

	// check if key exists
	value, ok := mpp["phone"]
	if ok {
		fmt.Println("ALL OK", value)
	} else {
		fmt.Println("key does not exist")
	}

	mpp2 := map[string]string{"name": "Gagan"}
	fmt.Println(maps.Equal(mpp, mpp2)) // true
}
