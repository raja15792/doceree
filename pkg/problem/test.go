package main

import (
	"fmt"
	"sort"
)

func main() {
	/*inputList := make([]string, 0)
	inputList = append(inputList, "3, 6")
	inputList = append(inputList, "2, 6")
	inputList = append(inputList, "9, 12")*/

	family := []struct {
		Name string
		Age  string
	}{
		{"Alice", "23"},
		{"David", "2"},
		{"Eve", "2"},
		{"Bob", "25"},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)
}
