package main

import "fmt"

func main() {
	// Two way to declare a empty map
	// Way 1:
	var mp1 = make(map[string]string)
	
	// Way 2:
	var mp2 map[string]string

	// -> different between way 1 and way 2
	fmt.Println(mp1 == nil)
	fmt.Println(mp2 == nil)
	// Result is false true

	
	// Update and add a element to the map
	// Add
	mp1["one"] = "Tran Quang Kiet"
	fmt.Println(mp1)
	// Update
	mp1["one"]= "Dang Quoc Hung"
	fmt.Println(mp1)


	// Map = nill can't add element
	mp2["one"] = "Nguyen Viet Hoang"
	fmt.Println(mp2)
	// Result is runtime error
	// Explain: when using the second way to declare a map, a nil map is generated. In Go, a nil map is a map that has not been initialized.


	// Delete a element from the map
	delete(mp1, "one")
	fmt.Println(mp1)


	// Check for specific elements in a Map
	mp1["one"] = "Tran Quang Kiet"
	name, ok := mp1["one"]
	if(name == "" && ok == false) {
		fmt.Println("Not found!")
	} else {
		fmt.Println("Found: ", name)
	}


	// Maps are references
	mp2 = mp1

	fmt.Println("Map 1 is: ", mp1)
	fmt.Println("Map 2 is: ", mp2)
	// Two mp1 and mp2 reference to "one" key
	// What happens if change value of mp1
	mp1["one"] = "Golang"
	fmt.Println("Map 1 after change is: ", mp1)
	fmt.Println("Map 2 after change is: ", mp2)


	// Iterate over Map
	for number,name := range mp1 {
		fmt.Printf("%s = %s\n", number, name)
	}


	// Maps in golang are unorder data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order. 
	var map1 = map[string]int {"one": 1, "two": 2, "three": 3}
	for key, value := range map1 {
		fmt.Printf("%s = %d\n", key, value)
	}
}