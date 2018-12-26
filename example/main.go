package main

import (
	"fmt"
	"sort"
)

func main() {
	mymap := map[string]int{
		"alice": 2, "aria": 8, "lucas": 4, "alex": 23,
	}

	keys := make([]string, 0)
	values := make([]int, 0)
	for k, v := range mymap {
		keys = append(keys, k)
		values = append(values, v)
	}

	sort.Strings(keys)
	sort.Ints(values)

	for _, e := range keys {
		fmt.Println(mymap[e])
	}
}
